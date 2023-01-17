package v1

const (
	//DataMountDir is the path where data volumes are mounted
	DataMountDir = "/mnt/data"
	//SecretsMountDir is the path where secrets are mounted
	SecretsMountDir = "/mnt/secrets"
	// LogicalClustersMountDir is the path where logical clusters resources are mounted
	LogicalClustersMountDir = "/mnt/logical_clusters"
	//SSLCertsMountDir is the path where ssl certs are mounted
	SSLCertsMountDir = "/mnt/sslcerts"
	//LogMountDir is the path where log volumes are mounted
	LogMountDir = "/mnt/log"
	//JMXServiceName is our standard k8s svc name for jmx
	JMXServiceName = "jmx"
	//JMXServicePort is our standard k8s svc port for jmx
	JMXServicePort = 7203
	//JolokiaServiceName is our standard k8s svc name for jolokia
	JolokiaServiceName = "jolokia"
	//JolokiaServicePort is our standard k8s svc port for jolokia
	JolokiaServicePort = 7777
	// PodComponentType is a pod label that states the component type of the pod
	PodComponentType = "type"
)
