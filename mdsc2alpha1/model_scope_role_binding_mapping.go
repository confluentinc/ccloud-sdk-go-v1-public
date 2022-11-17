package mdsv2alpha1

// ScopeRoleBindingMapping struct for ScopeRoleBindingMapping
type ScopeRoleBindingMapping struct {
	Scope        Scope                                   `json:"scope,omitempty"`
	Rolebindings map[string]map[string][]ResourcePattern `json:"rolebindings,omitempty"`
}
