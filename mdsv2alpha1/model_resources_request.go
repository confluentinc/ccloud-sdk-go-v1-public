package mdsv2alpha1

// ResourcesRequest struct for ResourcesRequest
type ResourcesRequest struct {
	Scope            Scope             `json:"scope"`
	ResourcePatterns []ResourcePattern `json:"resourcePatterns"`
}
