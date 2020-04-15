import Exception from './Exception'
import ServiceStatus from './Worktop/ServiceStatus'
import ServiceList from './Service/ServiceList'
import TaskTemplate from './Service/TaskTemplate'
import CreateTask from './Service/CreateTask'
import Istio from './Istio'

const routes = [
  {
    path: '/',
    exact: true,
    component: ServiceStatus
  },
  {
    path: '/worktop/overview',
    component: ServiceStatus
  },
  {
    path: '/service/serviceList',
    component: ServiceList
  },
  {
    path: '/service/taskTemplate',
    component: TaskTemplate
  },
  {
    path: '/service/createTask',
    component: CreateTask
  },
  {
    path: '/istio',
    component: Istio
  },
  {
    path: '/403',
    component: Exception
  },
  {
    path: '/404',
    component: Exception
  },
  {
    path: '/500',
    component: Exception
  }
]

export default routes
