package service

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/util"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	// ServiceInfo 处理在k8s中正常运行的service
	ServiceInfo *kubeInfo
	// IstioInfo 处理在k8s中运行的istio service
	IstioInfo *kubeInfo
)

type service struct {
	v1.Service      // 即k8s中定义的资源对象: service
	Pods       pods // 即k8s中定义的一组pod
}

type kubeInfo struct {
	mtx          *sync.RWMutex   // 读写锁
	wg           *sync.WaitGroup // WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。
	services     []service       // 一组service
	pods         []v1.Pod        // 一组pod
	namespaces   []v1.Namespace  // 一组namespace
	syncInterval time.Duration
	namespace    string
}

var (
	client     *kubernetes.Clientset
	kubeconfig string
)

/**
 * description: 初始化kube
 */
func InitKube() {
	// 初始化k8s客户端
	kubeconfig = util.Kubeconfig()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Error("[k8s] init client fail", "err", err)
		return
	}

	ServiceInfo = newKubeInfo("", time.Second*5)
	IstioInfo = newKubeInfo(bootstrap.Args.IstioNamespace, time.Second*5)

	// 周期性地从k8s与istio中同步集群信息
	go ServiceInfo.sync()
	go IstioInfo.sync()
}

func newKubeInfo(namespace string, syncInterval time.Duration) *kubeInfo {
	return &kubeInfo{
		mtx:          new(sync.RWMutex),
		wg:           new(sync.WaitGroup),
		services:     make([]service, 0),
		namespace:    namespace,
		syncInterval: syncInterval,
	}
}

type services []service
type namespaces []v1.Namespace

/**
 * description: 排除这些namespace下的service
 */
func (p services) Exclude(namespaces ...string) services {
	namespacesM := make(map[string]bool)
	for _, n := range namespaces {
		namespacesM[n] = true
	}

	retServices := make([]service, 0)
	for _, pod := range p {
		if _, ok := namespacesM[pod.Namespace]; !ok {
			retServices = append(retServices, pod)
		}
	}
	return retServices
}

/**
 * description: 返回所有/特定service
 */
func (k *kubeInfo) Services(uid string) services {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	if uid == "" {
		return k.services
	}

	ret := make([]service, 0)
	for _, s := range k.services {
		if string(s.UID) == uid {
			ret = append(ret, s)
			break
		}
	}
	return ret
}

/**
 * description: 返回所有/特定namespace
 */
func (k *kubeInfo) Namespaces(namespace string) namespaces {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	if namespace == "" {
		return k.namespaces
	}

	ret := make([]v1.Namespace, 0)
	for _, n := range k.namespaces {
		if string(n.Namespace) == namespace {
			ret = append(ret, n)
			break
		}
	}
	return ret
}

/**
 * description: 剔除指定的namespace
 */
func (n namespaces) Exclude(namespaces ...string) namespaces {
	namespacesM := make(map[string]bool)
	for _, n := range namespaces {
		namespacesM[n] = true
	}

	retNamespaces := make([]v1.Namespace, 0)
	for _, v := range n {
		if _, ok := namespacesM[v.Name]; !ok {
			retNamespaces = append(retNamespaces, v)
		}
	}
	return retNamespaces
}

/**
 * description: 定义service的简要信息结构体
 */
type KubeServiceStatus struct {
	UID        string `json:"uid"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ClusterIP  string `json:"clusterIP"`
	ExternalIP string `json:"externalIP"`
	Ports      string `json:"ports"`
	Age        string `json:"age"`
}

/**
 * description: 定义pod的简要信息结构体
 */
type KubePodStatus struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Ready    string `json:"ready"`
	Status   string `json:"status"` // Pending、Running、Succeeded、Failed、Unknown
	Restarts int    `json:"restarts"`
	Age      string `json:"age"`
}

/**
 * description: 返回pods的简要信息
 */
func (p pods) Status() []KubePodStatus {
	log.Info("[API] /api/diagnose Pods Status start", "ts", time.Now())

	pods := make([]KubePodStatus, 0, len(p))
	for _, item := range p {
		readyCnt, restartCnt, containerCnt := 0, 0, 0
		for _, c := range item.Status.ContainerStatuses {
			if c.Ready == true {
				readyCnt++
			}
			restartCnt += restartCnt
			containerCnt++
		}
		pods = append(pods, KubePodStatus{
			UID:      string(item.UID),
			Name:     item.Name,
			Ready:    fmt.Sprintf("%d/%d", readyCnt, containerCnt),
			Status:   string(item.Status.Phase),
			Restarts: readyCnt, // Todo
			Age:      time.Since(item.CreationTimestamp.Time).Truncate(time.Second).String(),
		})
	}

	log.Info("[API] /api/diagnose Pods Status end", "ts", time.Now())

	return pods
}

/**
 * description: 返回services的简要信息
 */
func (p services) Status() []KubeServiceStatus {
	log.Info("[API] /api/diagnose Services Status start", "ts", time.Now())

	components := make([]KubeServiceStatus, 0, len(p))
	for _, item := range p {
		ports := ""
		// 遍历每个service中的Ports
		for _, p := range item.Spec.Ports {
			ports += fmt.Sprintf(",%d/%s", p.Port, p.Protocol)
		}
		if ports != "" {
			ports = ports[1:]
		}
		components = append(components, KubeServiceStatus{
			UID:        string(item.UID),
			Name:       item.Name,
			Type:       string(item.Spec.Type),
			ClusterIP:  string(item.Spec.ClusterIP),
			ExternalIP: strings.Join(item.Spec.ExternalIPs, ","),
			Ports:      ports, // TODO
			Age:        time.Since(item.CreationTimestamp.Time).Truncate(time.Second).String(),
		})
	}
	log.Info("[API] /api/diagnose Services Status end", "ts", time.Now())

	return components
}

func (k *kubeInfo) podsFromK8S(labels map[string]string) pods {
	pods := make([]v1.Pod, 0)
	ls := ""
	if len(labels) != 0 {
		for k, v := range labels {
			ls += fmt.Sprintf(",%s=%s", k, v)
		}
		ls = ls[1:]
	}

	p, err := client.CoreV1().Pods(k.namespace).List(metav1.ListOptions{
		LabelSelector: ls,
	})
	if err != nil {
		log.Error("[k8s] get pods fail", "err", err)
		return pods
	}

	return p.Items
}

func (k *kubeInfo) Pods() pods {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	return k.pods
}

/**
 * description: 根据pod name返回一组pod
 */
func (k *kubeInfo) PodsByName(name string) pods {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	if name == "" {
		return k.pods
	}

	retPods := make([]v1.Pod, 0)
	for _, p := range k.pods {
		if p.Name == name {
			retPods = append(retPods, p)
		}
	}
	return retPods
}

type pods []v1.Pod // 即k8s中定义的一组pod

func (p pods) Exclude(namespaces ...string) pods {
	namespacesM := make(map[string]bool)
	for _, n := range namespaces {
		namespacesM[n] = true
	}

	retPods := make([]v1.Pod, 0)
	for _, pod := range p {
		if _, ok := namespacesM[pod.Namespace]; !ok {
			retPods = append(retPods, pod)
		}
	}
	return retPods
}

/**
 * description: k8s中service tree
 */
type Tree struct {
	Title         string `json:"title"`
	Key           string `json:"key"`
	GraphNodeName string `json:"graphNodeName"`
	Namespace     string `json:"namespace"`
	Children      []Tree `json:"children"`
}

/**
 * description: 返回service tree
 */
func (k *kubeInfo) Tree() []Tree {
	services := k.Services("").Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace)
	t := make([]Tree, 0, len(services))
	for _, s := range services {
		children := make([]Tree, 0, len(s.Pods))
		for _, pod := range s.Pods {
			children = append(children, Tree{
				Title:         pod.Name,
				Key:           string(pod.UID),
				Namespace:     pod.Namespace,
				GraphNodeName: fmt.Sprintf("%s-%s", s.Name, pod.Labels["version"]),
			})
		}
		t = append(t, Tree{
			Title:     s.Name,
			Key:       string(s.UID),
			Namespace: s.Namespace,
			Children:  children,
		})
	}
	return t
}

/**
 * description: 从k8s和istio中周期性同步集群信息
 */
func (k *kubeInfo) sync() {
	for {
		log.Debug("[Kube] sync start", "svcs", len(k.services), "namespace", k.namespace, "time", time.Now())
		svcs, err := client.CoreV1().Services(k.namespace).List(metav1.ListOptions{
			LabelSelector: "provider!=kubernetes",
		})
		if err != nil {
			// panic(err.Error())
			log.Error("[k8s] get services err", "err", err)
		}
		ns, err := client.CoreV1().Namespaces().List(metav1.ListOptions{
			LabelSelector: "provider!=kubernetes",
		})
		if err != nil {
			// panic(err.Error())
			log.Error("[k8s] get namespaces err", "err", err)
		}

		// get services' and pods' data from Kubernetes
		var serviceCh = make(chan service, 200)
		k.wg.Add(len(svcs.Items))
		for _, i := range svcs.Items {
			go func(i v1.Service) {
				s := service{}
				s.Service = i
				s.Pods = k.podsFromK8S(i.Spec.Selector)
				serviceCh <- s
			}(i)
		}

		services := make([]service, 0, len(svcs.Items))
		tmpPods := make(map[string]v1.Pod)
		pods := make(pods, 0)
		go func() {
			for s := range serviceCh {
				services = append(services, s)
				for _, p := range s.Pods {
					if _, ok := tmpPods[string(p.UID)]; !ok {
						pods = append(pods, p)
					}
				}
				k.wg.Done()
			}
		}()
		k.wg.Wait()
		close(serviceCh)

		k.mtx.Lock()
		k.services = services
		k.namespaces = ns.Items
		k.pods = pods
		k.mtx.Unlock()

		log.Debug("[Kube] sync end", "svcs", len(k.services), "namespace", k.namespace, "time", time.Now())
		time.Sleep(k.syncInterval)
	}
}
