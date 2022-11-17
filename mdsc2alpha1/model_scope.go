package mdsv2alpha1

// Scope struct for Scope
type Scope struct {
	Path     []string      `json:"path,omitempty"`
	Clusters ScopeClusters `json:"clusters,omitempty"`
}
