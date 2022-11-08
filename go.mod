module github.com/confluentinc/ccloud-sdk-go-v1-internal

require (
	github.com/confluentinc/proto-go-setter v0.0.0-20201026155413-c6ceb267ee65
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace (
	github.com/confluentinc/protoc-gen-ccloud => github.com/confluentinc/protoc-gen-ccloud v0.0.3
	github.com/influxdata/influxdb1-client => github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	k8s.io/api => k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
)

go 1.16
