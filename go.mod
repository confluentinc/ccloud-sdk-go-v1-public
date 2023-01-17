module github.com/confluentinc/ccloud-sdk-go-v1-public

require (
	github.com/confluentinc/proto-go-setter v0.3.0
	github.com/dghubble/sling v1.4.1
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	github.com/gogo/googleapis v1.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/travisjeffery/proto-go-sql v0.0.0-20190911121832-39ff47280e87
	github.com/ugorji/go/codec v1.2.8
	golang.org/x/oauth2 v0.4.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0 // indirect
	k8s.io/klog v1.0.0 // indirect
)

replace (
	github.com/confluentinc/protoc-gen-ccloud => github.com/confluentinc/protoc-gen-ccloud v0.0.3
	github.com/influxdata/influxdb1-client => github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	k8s.io/api => k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => k8s.io/apiserver v0.17.0
)

go 1.16
