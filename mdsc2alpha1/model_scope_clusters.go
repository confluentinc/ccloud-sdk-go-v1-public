package mdsv2alpha1

// ScopeClusters struct for ScopeClusters
type ScopeClusters struct {
	KafkaCluster          string `json:"kafka-cluster,omitempty"`
	ConnectCluster        string `json:"connect-cluster,omitempty"`
	KsqlCluster           string `json:"ksql-cluster,omitempty"`
	SchemaRegistryCluster string `json:"schema-registry-cluster,omitempty"`
}
