module github.com/confluentinc/ccloud-sdk-go-v1-public

go 1.23.8

require (
	github.com/confluentinc/proto-go-setter v0.3.0
	github.com/dghubble/sling v1.4.1
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	github.com/gogo/googleapis v1.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/travisjeffery/mocker v1.1.0
	github.com/travisjeffery/proto-go-sql v0.0.0-20190911121832-39ff47280e87
	github.com/ugorji/go/codec v1.2.8
	golang.org/x/oauth2 v0.4.0
	k8s.io/api v0.17.0
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/golang/mock v1.4.4 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/lyft/protoc-gen-star v0.6.1 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/inf.v0 v0.9.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
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
