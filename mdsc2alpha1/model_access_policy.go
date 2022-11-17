package mdsv2alpha1

// AccessPolicy struct for AccessPolicy
type AccessPolicy struct {
	BindingScope      string      `json:"bindingScope,omitempty"`
	BindWithResource  bool        `json:"bindWithResource"`
	AllowedOperations []Operation `json:"allowedOperations,omitempty"`
}
