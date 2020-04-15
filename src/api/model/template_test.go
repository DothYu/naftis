package model

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

var (
	requestRouting = map[string]string{
		"DestinationHost":   "rating",
		"DestinationSubset": "v1",
	}
)

func TestRequestRouting(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: {{.DestinationHost}}
  ...
spec:
  hosts:
  - details
  http:
  - route:
    - destination:
        host: {{.DestinationHost}}
        subset: {{.DestinationSubset}}`)
	assert.NoError(t, err)

	var b bytes.Buffer
	err = tmpl.Execute(&b, requestRouting)
	assert.NoError(t, err)

	actual := b.String()
	expect := `
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: rating
  ...
spec:
  hosts:
  - details
  http:
  - route:
    - destination:
        host: rating
        subset: v1`

	assert.Equal(t, expect, actual)
}
