package mdsv2alpha1

// ResourcePattern struct for ResourcePattern
type ResourcePattern struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	PatternType  string `json:"patternType"`
}
