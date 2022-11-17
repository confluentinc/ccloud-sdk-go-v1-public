package mdsv2alpha1

// Operation struct for Operation
type Operation struct {
	ResourceType string   `json:"resourceType,omitempty"`
	Operations   []string `json:"operations,omitempty"`
}
