module github.com/confluentinc/ccloud-sdk-go-v1-public

require (
	github.com/confluentinc/cc-structs/common v0.1373.0 // indirect
	github.com/confluentinc/cc-structs/kafka/core v0.1373.0
	github.com/confluentinc/cc-structs/kafka/org v0.1373.0
	github.com/confluentinc/cc-structs/kafka/product/core v0.1373.0 // indirect
	github.com/confluentinc/cc-structs/kafka/scheduler v0.1373.0 // indirect
	github.com/confluentinc/cc-structs/kafka/util v0.1373.0 // indirect
	github.com/confluentinc/cc-structs/operator v0.1373.0 // indirect
	github.com/confluentinc/proto-go-setter v0.0.0-20201026155413-c6ceb267ee65
	github.com/dghubble/sling v1.2.1-0.20181125223409-7458fd7fa70b
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.8.0
	github.com/travisjeffery/proto-go-sql v0.0.0-20190911121832-39ff47280e87
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b
	golang.org/x/sys v0.0.0-20220224120231-95c6836cb0e7 // indirect
	google.golang.org/genproto v0.0.0-20220222213610-43724f9ea8cf // indirect
)

replace (
	github.com/confluentinc/protoc-gen-ccloud => github.com/confluentinc/protoc-gen-ccloud v0.0.3
	github.com/influxdata/influxdb1-client => github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	k8s.io/api => k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
)

go 1.16
