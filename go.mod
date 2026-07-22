module github.com/confluentinc/ccloud-sdk-go-v1-public

go 1.25.0

require (
	github.com/confluentinc/proto-go-setter v0.3.0
	github.com/dghubble/sling v1.4.1
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/gogo/googleapis v1.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.4
	github.com/travisjeffery/mocker v1.1.0
	github.com/travisjeffery/proto-go-sql v0.0.0-20190911121832-39ff47280e87
	github.com/ugorji/go/codec v1.2.8
	golang.org/x/oauth2 v0.36.0
	k8s.io/api v0.17.0
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lyft/protoc-gen-star/v2 v2.0.4 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/mod v0.34.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	golang.org/x/tools v0.43.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/inf.v0 v0.9.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.17.0 // indirect
	k8s.io/klog v1.0.0 // indirect
)

replace (
	github.com/confluentinc/protoc-gen-ccloud => github.com/confluentinc/protoc-gen-ccloud v0.0.3
	github.com/influxdata/influxdb1-client => github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	google.golang.org/grpc => google.golang.org/grpc v1.82.1
	k8s.io/api => k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => k8s.io/apiserver v0.17.0
)
