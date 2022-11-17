package mdsv2alpha1

// Role struct for Role
type Role struct {
	Name     string         `json:"name,omitempty"`
	Policies []AccessPolicy `json:"policies,omitempty"`
}
