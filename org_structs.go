package ccloud

import (
	_ "github.com/confluentinc/proto-go-setter"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/travisjeffery/proto-go-sql"
)

type Organization struct {
	Id               int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	Name             string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`
	Deactivated      bool             `protobuf:"varint,3,opt,name=deactivated,proto3" json:"deactivated,omitempty" db:"deactivated,omitempty" url:"deactivated,omitempty"`
	StripeCustomerId string           `protobuf:"bytes,4,opt,name=stripe_customer_id,json=stripeCustomerId,proto3" json:"stripe_customer_id,omitempty" db:"stripe_customer_id,omitempty" url:"stripe_customer_id,omitempty"`
	Created          *types.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty" db:"created,omitempty" url:"created,omitempty"`
	Modified         *types.Timestamp `protobuf:"bytes,6,opt,name=modified,proto3" json:"modified,omitempty" db:"modified,omitempty" url:"modified,omitempty"`
	BillingEmail     string           `protobuf:"bytes,7,opt,name=billing_email,json=billingEmail,proto3" json:"billing_email,omitempty" db:"billing_email,omitempty" url:"billing_email,omitempty"`
	Plan             *Plan            `protobuf:"bytes,8,opt,name=plan,proto3" json:"plan,omitempty" db:"plan,omitempty" url:"plan,omitempty"`
	Saml             *Saml            `protobuf:"bytes,9,opt,name=saml,proto3" json:"saml,omitempty" db:"saml,omitempty" url:"saml,omitempty"`
	Sso              *Sso             `protobuf:"bytes,10,opt,name=sso,proto3" json:"sso,omitempty" db:"sso,omitempty" url:"sso,omitempty"`
	Marketplace      *Marketplace     `protobuf:"bytes,11,opt,name=marketplace,proto3" json:"marketplace,omitempty" db:"marketplace,omitempty" url:"marketplace,omitempty"`
	ResourceId       string           `protobuf:"bytes,12,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty" db:"resource_id,omitempty" url:"resource_id,omitempty"`
	//derived flags
	HasEntitlement bool      `protobuf:"varint,13,opt,name=has_entitlement,json=hasEntitlement,proto3" json:"has_entitlement,omitempty" db:"has_entitlement,omitempty" url:"has_entitlement,omitempty"`
	ShowBilling    bool      `protobuf:"varint,14,opt,name=show_billing,json=showBilling,proto3" json:"show_billing,omitempty" db:"show_billing,omitempty" url:"show_billing,omitempty"`
	AuditLog       *AuditLog `protobuf:"bytes,15,opt,name=audit_log,json=auditLog,proto3" json:"audit_log,omitempty" db:"audit_log,omitempty" url:"audit_log,omitempty"`
	// Derived flag from active orders (based on the current time), indicating if an org has commitment with us.
	HasCommitment bool `protobuf:"varint,16,opt,name=has_commitment,json=hasCommitment,proto3" json:"has_commitment,omitempty" db:"has_commitment,omitempty" url:"has_commitment,omitempty"` // Deprecated: Do not use.
	//derived marketplace subscription
	MarketplaceSubscription MarketplaceSubscription `protobuf:"varint,17,opt,name=marketplace_subscription,json=marketplaceSubscription,proto3,enum=kafka.org.v1.MarketplaceSubscription" json:"marketplace_subscription,omitempty" db:"marketplace_subscription,omitempty" url:"marketplace_subscription,omitempty"`
	DeactivatedAt           *types.Timestamp        `protobuf:"bytes,18,opt,name=deactivated_at,json=deactivatedAt,proto3" json:"deactivated_at,omitempty" db:"deactivated_at,omitempty" url:"deactivated_at,omitempty"`
	SuspensionStatus        *SuspensionStatus       `protobuf:"bytes,19,opt,name=suspension_status,json=suspensionStatus,proto3" json:"suspension_status,omitempty" db:"suspension_status,omitempty" url:"suspension_status,omitempty"`
	DisplayLabel            string                  `protobuf:"bytes,20,opt,name=display_label,json=displayLabel,proto3" json:"display_label,omitempty" db:"display_label,omitempty" url:"display_label,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                `json:"-"`
	XXX_unrecognized        []byte                  `json:"-"`
	XXX_sizecache           int32                   `json:"-"`
}
