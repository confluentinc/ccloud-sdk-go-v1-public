package ccloud

import (
	encoding_binary "encoding/binary"
	"fmt"
	io "io"
	math "math"

	_ "github.com/confluentinc/proto-go-setter"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/travisjeffery/proto-go-sql"
)

type User struct {
	Id                   int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" url:"id" db:"id,omitempty"`
	Email                string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" url:"email" db:"email,omitempty"`
	FirstName            string            `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty" url:"first_name" db:"first_name,omitempty"`
	LastName             string            `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty" url:"last_name" db:"last_name,omitempty"`
	OrganizationId       int32             `protobuf:"varint,6,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty" url:"organization_id" db:"organization_id,omitempty"`
	Deactivated          bool              `protobuf:"varint,7,opt,name=deactivated,proto3" json:"deactivated,omitempty" db:"deactivated,omitempty" url:"deactivated,omitempty"`
	Verified             *types.Timestamp  `protobuf:"bytes,11,opt,name=verified,proto3" json:"verified,omitempty" db:"verified,omitempty" url:"verified,omitempty"`
	Created              *types.Timestamp  `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty" db:"created,omitempty" url:"created,omitempty"`
	Modified             *types.Timestamp  `protobuf:"bytes,9,opt,name=modified,proto3" json:"modified,omitempty" db:"modified,omitempty" url:"modified,omitempty"`
	ServiceName          string            `protobuf:"bytes,12,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty" db:"service_name,omitempty" url:"service_name,omitempty"`
	ServiceDescription   string            `protobuf:"bytes,13,opt,name=service_description,json=serviceDescription,proto3" json:"service_description,omitempty" db:"service_description,omitempty" url:"service_description,omitempty"`
	ServiceAccount       bool              `protobuf:"varint,14,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty" db:"service_account,omitempty" url:"service_account,omitempty"`
	Preferences          map[string]string `protobuf:"bytes,16,rep,name=preferences,proto3" json:"preferences,omitempty" db:"preferences,omitempty" url:"preferences,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Internal             bool              `protobuf:"varint,17,opt,name=internal,proto3" json:"internal,omitempty" db:"internal,omitempty" url:"internal,omitempty"`
	ResourceId           string            `protobuf:"bytes,18,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty" db:"resource_id,omitempty" url:"resource_id,omitempty"`
	DeactivatedAt        *types.Timestamp  `protobuf:"bytes,19,opt,name=deactivated_at,json=deactivatedAt,proto3" json:"deactivated_at,omitempty" db:"deactivated_at,omitempty" url:"deactivated_at,omitempty"`
	SocialConnection     string            `protobuf:"bytes,20,opt,name=social_connection,json=socialConnection,proto3" json:"social_connection,omitempty" db:"social_connection,omitempty" url:"social_connection,omitempty"`
	AuthType             AuthType          `protobuf:"varint,21,opt,name=auth_type,json=authType,proto3,enum=kafka.org.v1.AuthType" json:"auth_type,omitempty" db:"auth_type,omitempty" url:"auth_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetSocialConnection() string {
	if m != nil {
		return m.SocialConnection
	}
	return ""
}

func (m *User) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_AUTH_TYPE_UNKNOWN
}

type AuthType int32

const (
	AuthType_AUTH_TYPE_UNKNOWN AuthType = 0
	AuthType_AUTH_TYPE_LOCAL   AuthType = 1
	AuthType_AUTH_TYPE_SSO     AuthType = 2
)

var AuthType_name = map[int32]string{
	0: "AUTH_TYPE_UNKNOWN",
	1: "AUTH_TYPE_LOCAL",
	2: "AUTH_TYPE_SSO",
}

var AuthType_value = map[string]int32{
	"AUTH_TYPE_UNKNOWN": 0,
	"AUTH_TYPE_LOCAL":   1,
	"AUTH_TYPE_SSO":     2,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}

type Account struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`
	OrganizationId       int32            `protobuf:"varint,4,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty" db:"organization_id,omitempty" url:"organization_id,omitempty"`
	Deactivated          bool             `protobuf:"varint,5,opt,name=deactivated,proto3" json:"deactivated,omitempty" db:"deactivated,omitempty" url:"deactivated,omitempty"`
	Created              *types.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty" db:"created,omitempty" url:"created,omitempty"`
	Modified             *types.Timestamp `protobuf:"bytes,7,opt,name=modified,proto3" json:"modified,omitempty" db:"modified,omitempty" url:"modified,omitempty"`
	Config               *AccountConfig   `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty" db:"config,omitempty" url:"config,omitempty"`
	Internal             bool             `protobuf:"varint,9,opt,name=internal,proto3" json:"internal,omitempty" db:"internal,omitempty" url:"internal,omitempty"`
	DeactivatedAt        *types.Timestamp `protobuf:"bytes,10,opt,name=deactivated_at,json=deactivatedAt,proto3" json:"deactivated_at,omitempty" db:"deactivated_at,omitempty" url:"deactivated_at,omitempty"`
	OrgResourceId        string           `protobuf:"bytes,11,opt,name=org_resource_id,json=orgResourceId,proto3" json:"org_resource_id,omitempty" db:"org_resource_id,omitempty" url:"org_resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AccountConfig struct {
	MaxKafkaClusters     int32    `protobuf:"varint,1,opt,name=max_kafka_clusters,json=maxKafkaClusters,proto3" json:"max_kafka_clusters,omitempty" db:"max_kafka_clusters,omitempty" url:"max_kafka_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountConfig) Reset()         { *m = AccountConfig{} }
func (m *AccountConfig) String() string { return proto.CompactTextString(m) }
func (*AccountConfig) ProtoMessage()    {}

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

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Organization) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Organization) GetSuspensionStatus() *SuspensionStatus {
	if m != nil {
		return m.SuspensionStatus
	}
	return nil
}

func (m *Organization) GetAuditLog() *AuditLog {
	if m != nil {
		return m.AuditLog
	}
	return nil
}

func (m *Organization) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

func (m *Organization) GetDeactivated() bool {
	if m != nil {
		return m.Deactivated
	}
	return false
}

// Represents the customer-accessible audit log cluster info for the organization
type AuditLog struct {
	ClusterId                string   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" db:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	AccountId                string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty" db:"account_id,omitempty" url:"account_id,omitempty"`
	ServiceAccountId         int32    `protobuf:"varint,3,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty" db:"service_account_id,omitempty" url:"service_account_id,omitempty"`
	TopicName                string   `protobuf:"bytes,4,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty" db:"topic_name,omitempty" url:"topic_name,omitempty"`
	ServiceAccountResourceId string   `protobuf:"bytes,5,opt,name=service_account_resource_id,json=serviceAccountResourceId,proto3" json:"service_account_resource_id,omitempty" db:"service_account_resource_id,omitempty" url:"service_account_resource_id,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}

func (m *AuditLog) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *AuditLog) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AuditLog) GetServiceAccountId() int32 {
	if m != nil {
		return m.ServiceAccountId
	}
	return 0
}

func (m *AuditLog) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *AuditLog) GetServiceAccountResourceId() string {
	if m != nil {
		return m.ServiceAccountResourceId
	}
	return ""
}

type SuspensionStatus struct {
	Suspended              *types.Timestamp     `protobuf:"bytes,1,opt,name=suspended,proto3" json:"suspended,omitempty" db:"suspended,omitempty" url:"suspended,omitempty"`
	Status                 SuspensionStatusType `protobuf:"varint,2,opt,name=status,proto3,enum=kafka.org.v1.SuspensionStatusType" json:"status,omitempty" db:"status,omitempty" url:"status,omitempty"`
	EventType              SuspensionEventType  `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=kafka.org.v1.SuspensionEventType" json:"event_type,omitempty" db:"event_type,omitempty" url:"event_type,omitempty"`
	ScheduledDeactivatedAt *types.Timestamp     `protobuf:"bytes,4,opt,name=scheduled_deactivated_at,json=scheduledDeactivatedAt,proto3" json:"scheduled_deactivated_at,omitempty" db:"scheduled_deactivated_at,omitempty" url:"scheduled_deactivated_at,omitempty"`
	Version                int64                `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty" db:"version,omitempty" url:"version,omitempty"`
	ErrorMessage           string               `protobuf:"bytes,6,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty" db:"error_message,omitempty" url:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *SuspensionStatus) GetStatus() SuspensionStatusType {
	if m != nil {
		return m.Status
	}
	return SuspensionStatusType_SUSPENSION_UNKNOWN
}

func (m *SuspensionStatus) GetEventType() SuspensionEventType {
	if m != nil {
		return m.EventType
	}
	return SuspensionEventType_SUSPENSION_EVENT_UNKNOWN
}

type Plan struct {
	TaxAddress           *Address         `protobuf:"bytes,1,opt,name=tax_address,json=taxAddress,proto3" json:"tax_address,omitempty" db:"tax_address,omitempty" url:"tax_address,omitempty"`
	ProductLevel         ProductLevel     `protobuf:"varint,2,opt,name=product_level,json=productLevel,proto3,enum=kafka.org.v1.ProductLevel" json:"product_level,omitempty" db:"product_level,omitempty" url:"product_level,omitempty"`
	TrialStart           *types.Timestamp `protobuf:"bytes,3,opt,name=trial_start,json=trialStart,proto3" json:"trial_start,omitempty" db:"trial_start,omitempty" url:"trial_start,omitempty"`
	TrialEnd             *types.Timestamp `protobuf:"bytes,4,opt,name=trial_end,json=trialEnd,proto3" json:"trial_end,omitempty" db:"trial_end,omitempty" url:"trial_end,omitempty"`
	PlanStart            *types.Timestamp `protobuf:"bytes,5,opt,name=plan_start,json=planStart,proto3" json:"plan_start,omitempty" db:"plan_start,omitempty" url:"plan_start,omitempty"`
	PlanEnd              *types.Timestamp `protobuf:"bytes,6,opt,name=plan_end,json=planEnd,proto3" json:"plan_end,omitempty" db:"plan_end,omitempty" url:"plan_end,omitempty"`
	Product              *Plan_Product    `protobuf:"bytes,8,opt,name=product,proto3" json:"product,omitempty" db:"product,omitempty" url:"product,omitempty"`
	Billing              *Plan_Billing    `protobuf:"bytes,9,opt,name=billing,proto3" json:"billing,omitempty" db:"billing,omitempty" url:"billing,omitempty"`
	ReferralCode         string           `protobuf:"bytes,10,opt,name=referral_code,json=referralCode,proto3" json:"referral_code,omitempty" db:"referral_code,omitempty" url:"referral_code,omitempty"`
	AcceptTos            *types.BoolValue `protobuf:"bytes,11,opt,name=accept_tos,json=acceptTos,proto3" json:"accept_tos,omitempty" db:"accept_tos,omitempty" url:"accept_tos,omitempty"`
	AllowMultiTenant     bool             `protobuf:"varint,12,opt,name=allow_multi_tenant,json=allowMultiTenant,proto3" json:"allow_multi_tenant,omitempty" db:"allow_multi_tenant,omitempty" url:"allow_multi_tenant,omitempty"`
	AcceptTosPlatform    *types.BoolValue `protobuf:"bytes,13,opt,name=accept_tos_platform,json=acceptTosPlatform,proto3" json:"accept_tos_platform,omitempty" db:"accept_tos_platform,omitempty" url:"accept_tos_platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Plan) GetBilling() *Plan_Billing {
	if m != nil {
		return m.Billing
	}
	return nil
}

// representation of SSO configs on the DB records themselves
type Sso struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" db:"enabled,omitempty" url:"enabled,omitempty"`
	Auth0ConnectionName  string   `protobuf:"bytes,2,opt,name=auth0_connection_name,json=auth0ConnectionName,proto3" json:"auth0_connection_name,omitempty" db:"auth0_connection_name,omitempty" url:"auth0_connection_name,omitempty"`
	TenantId             string   `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty" db:"tenant_id,omitempty" url:"tenant_id,omitempty"`
	MultiTenant          bool     `protobuf:"varint,4,opt,name=multi_tenant,json=multiTenant,proto3" json:"multi_tenant,omitempty" db:"multi_tenant,omitempty" url:"multi_tenant,omitempty"`
	Mode                 SsoMode  `protobuf:"varint,7,opt,name=mode,proto3,enum=kafka.org.v1.SsoMode" json:"mode,omitempty" db:"mode,omitempty" url:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type Marketplace struct {
	Partner              MarketplacePartner       `protobuf:"varint,1,opt,name=partner,proto3,enum=kafka.core.v1.MarketplacePartner" json:"partner,omitempty" db:"partner,omitempty" url:"partner,omitempty"`
	CustomerId           string                   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty" db:"customer_id,omitempty" url:"customer_id,omitempty"`
	CustomerState        MarketplaceCustomerState `protobuf:"varint,3,opt,name=customer_state,json=customerState,proto3,enum=kafka.org.v1.MarketplaceCustomerState" json:"customer_state,omitempty" db:"customer_state,omitempty" url:"customer_state,omitempty"`
	ConsoleIntegrated    bool                     `protobuf:"varint,4,opt,name=console_integrated,json=consoleIntegrated,proto3" json:"console_integrated,omitempty" db:"console_integrated,omitempty" url:"console_integrated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

type Address struct {
	Street1              string   `protobuf:"bytes,1,opt,name=street1,proto3" json:"street1,omitempty" db:"street1,omitempty" url:"street1,omitempty"`
	Street2              string   `protobuf:"bytes,2,opt,name=street2,proto3" json:"street2,omitempty" db:"street2,omitempty" url:"street2,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty" db:"city,omitempty" url:"city,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty" db:"state,omitempty" url:"state,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty" db:"country,omitempty" url:"country,omitempty"`
	Zip                  string   `protobuf:"bytes,6,opt,name=zip,proto3" json:"zip,omitempty" db:"zip,omitempty" url:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type MarketplaceCustomerState int32

const (
	MarketplaceCustomerState_CUSTOMER_APPROVAL_PENDING MarketplaceCustomerState = 0
	MarketplaceCustomerState_CUSTOMER_APPROVAL_SENT    MarketplaceCustomerState = 1
	MarketplaceCustomerState_CUSTOMER_APPROVED         MarketplaceCustomerState = 2
	MarketplaceCustomerState_CUSTOMER_DELETED          MarketplaceCustomerState = 3
)

var MarketplaceCustomerState_name = map[int32]string{
	0: "CUSTOMER_APPROVAL_PENDING",
	1: "CUSTOMER_APPROVAL_SENT",
	2: "CUSTOMER_APPROVED",
	3: "CUSTOMER_DELETED",
}

var MarketplaceCustomerState_value = map[string]int32{
	"CUSTOMER_APPROVAL_PENDING": 0,
	"CUSTOMER_APPROVAL_SENT":    1,
	"CUSTOMER_APPROVED":         2,
	"CUSTOMER_DELETED":          3,
}

func (x MarketplaceCustomerState) String() string {
	return proto.EnumName(MarketplaceCustomerState_name, int32(x))
}

type SuspensionEventType int32

const (
	SuspensionEventType_SUSPENSION_EVENT_UNKNOWN                              SuspensionEventType = 0
	SuspensionEventType_SUSPENSION_EVENT_END_OF_FREE_TRIAL                    SuspensionEventType = 1
	SuspensionEventType_SUSPENSION_EVENT_CUSTOMER_INITIATED_ORG_DEACTIVATION  SuspensionEventType = 2
	SuspensionEventType_SUSPENSION_EVENT_MARKETPLACE_ENTITLEMENT_CANCELLATION SuspensionEventType = 3
	SuspensionEventType_SUSPENSION_EVENT_NO_PAYMENT                           SuspensionEventType = 4
	SuspensionEventType_SUSPENSION_EVENT_SECURITY_INCIDENT_OR_ABUSE           SuspensionEventType = 5
	SuspensionEventType_SUSPENSION_EVENT_OTHERS                               SuspensionEventType = 6
	SuspensionEventType_SUSPENSION_EVENT_INTERNAL_INITIATED_ORG_DEACTIVATION  SuspensionEventType = 7
)

var SuspensionEventType_name = map[int32]string{
	0: "SUSPENSION_EVENT_UNKNOWN",
	1: "SUSPENSION_EVENT_END_OF_FREE_TRIAL",
	2: "SUSPENSION_EVENT_CUSTOMER_INITIATED_ORG_DEACTIVATION",
	3: "SUSPENSION_EVENT_MARKETPLACE_ENTITLEMENT_CANCELLATION",
	4: "SUSPENSION_EVENT_NO_PAYMENT",
	5: "SUSPENSION_EVENT_SECURITY_INCIDENT_OR_ABUSE",
	6: "SUSPENSION_EVENT_OTHERS",
	7: "SUSPENSION_EVENT_INTERNAL_INITIATED_ORG_DEACTIVATION",
}

var SuspensionEventType_value = map[string]int32{
	"SUSPENSION_EVENT_UNKNOWN":                              0,
	"SUSPENSION_EVENT_END_OF_FREE_TRIAL":                    1,
	"SUSPENSION_EVENT_CUSTOMER_INITIATED_ORG_DEACTIVATION":  2,
	"SUSPENSION_EVENT_MARKETPLACE_ENTITLEMENT_CANCELLATION": 3,
	"SUSPENSION_EVENT_NO_PAYMENT":                           4,
	"SUSPENSION_EVENT_SECURITY_INCIDENT_OR_ABUSE":           5,
	"SUSPENSION_EVENT_OTHERS":                               6,
	"SUSPENSION_EVENT_INTERNAL_INITIATED_ORG_DEACTIVATION":  7,
}

func (x SuspensionEventType) String() string {
	return proto.EnumName(SuspensionEventType_name, int32(x))
}

type SuspensionStatusType int32

const (
	SuspensionStatusType_SUSPENSION_UNKNOWN       SuspensionStatusType = 0
	SuspensionStatusType_SUSPENSION_IN_PROGRESS   SuspensionStatusType = 1
	SuspensionStatusType_SUSPENSION_COMPLETED     SuspensionStatusType = 2
	SuspensionStatusType_UNSUSPENSION_IN_PROGRESS SuspensionStatusType = 3
	SuspensionStatusType_SUSPENSION_FAILED        SuspensionStatusType = 4
	SuspensionStatusType_UNSUSPENSION_FAILED      SuspensionStatusType = 5
)

var SuspensionStatusType_name = map[int32]string{
	0: "SUSPENSION_UNKNOWN",
	1: "SUSPENSION_IN_PROGRESS",
	2: "SUSPENSION_COMPLETED",
	3: "UNSUSPENSION_IN_PROGRESS",
	4: "SUSPENSION_FAILED",
	5: "UNSUSPENSION_FAILED",
}

var SuspensionStatusType_value = map[string]int32{
	"SUSPENSION_UNKNOWN":       0,
	"SUSPENSION_IN_PROGRESS":   1,
	"SUSPENSION_COMPLETED":     2,
	"UNSUSPENSION_IN_PROGRESS": 3,
	"SUSPENSION_FAILED":        4,
	"UNSUSPENSION_FAILED":      5,
}

func (x SuspensionStatusType) String() string {
	return proto.EnumName(SuspensionStatusType_name, int32(x))
}

type MarketplaceSubscription int32

const (
	MarketplaceSubscription_SUBSCRIPTION_NONE    MarketplaceSubscription = 0
	MarketplaceSubscription_SUBSCRIPTION_PENDING MarketplaceSubscription = 1
	MarketplaceSubscription_SUBSCRIPTION_ACTIVE  MarketplaceSubscription = 2
)

var MarketplaceSubscription_name = map[int32]string{
	0: "SUBSCRIPTION_NONE",
	1: "SUBSCRIPTION_PENDING",
	2: "SUBSCRIPTION_ACTIVE",
}

var MarketplaceSubscription_value = map[string]int32{
	"SUBSCRIPTION_NONE":    0,
	"SUBSCRIPTION_PENDING": 1,
	"SUBSCRIPTION_ACTIVE":  2,
}

func (x MarketplaceSubscription) String() string {
	return proto.EnumName(MarketplaceSubscription_name, int32(x))
}

type SsoMode int32

const (
	SsoMode_SSO_MODE_UNKNOWN    SsoMode = 0
	SsoMode_SSO_MODE_RESTRICTED SsoMode = 1
	SsoMode_SSO_MODE_LAX        SsoMode = 2
)

var SsoMode_name = map[int32]string{
	0: "SSO_MODE_UNKNOWN",
	1: "SSO_MODE_RESTRICTED",
	2: "SSO_MODE_LAX",
}

var SsoMode_value = map[string]int32{
	"SSO_MODE_UNKNOWN":    0,
	"SSO_MODE_RESTRICTED": 1,
	"SSO_MODE_LAX":        2,
}

func (x SsoMode) String() string {
	return proto.EnumName(SsoMode_name, int32(x))
}

type ProductLevel int32

const (
	ProductLevel_DEVELOPER  ProductLevel = 0
	ProductLevel_TEAM       ProductLevel = 1
	ProductLevel_ENTERPRISE ProductLevel = 2
	ProductLevel_UNIFIED    ProductLevel = 3
)

var ProductLevel_name = map[int32]string{
	0: "DEVELOPER",
	1: "TEAM",
	2: "ENTERPRISE",
	3: "UNIFIED",
}

var ProductLevel_value = map[string]int32{
	"DEVELOPER":  0,
	"TEAM":       1,
	"ENTERPRISE": 2,
	"UNIFIED":    3,
}

func (x ProductLevel) String() string {
	return proto.EnumName(ProductLevel_name, int32(x))
}

type Plan_Product struct {
	ReadThroughputMb  float64 `protobuf:"fixed64,1,opt,name=read_throughput_mb,json=readThroughputMb,proto3" json:"read_throughput_mb,omitempty" db:"read_throughput_mb,omitempty" url:"read_throughput_mb,omitempty"`
	WriteThroughputMb float64 `protobuf:"fixed64,2,opt,name=write_throughput_mb,json=writeThroughputMb,proto3" json:"write_throughput_mb,omitempty" db:"write_throughput_mb,omitempty" url:"write_throughput_mb,omitempty"`
	// max data that can be retained, usable, not counting total allocated for replication
	StorageGb            int32    `protobuf:"varint,3,opt,name=storage_gb,json=storageGb,proto3" json:"storage_gb,omitempty" db:"storage_gb,omitempty" url:"storage_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type Plan_Billing struct {
	Method   BillingMethod   `protobuf:"varint,1,opt,name=method,proto3,enum=kafka.org.v1.BillingMethod" json:"method,omitempty" db:"method,omitempty" url:"method,omitempty"`
	Interval BillingInterval `protobuf:"varint,2,opt,name=interval,proto3,enum=kafka.org.v1.BillingInterval" json:"interval,omitempty" db:"interval,omitempty" url:"interval,omitempty"`
	// in hundredths of a cent, aggregates the running cost of all clusters
	// in the current billing cycle
	AccruedThisCycle     uint64   `protobuf:"varint,3,opt,name=accrued_this_cycle,json=accruedThisCycle,proto3" json:"accrued_this_cycle,omitempty" db:"accrued_this_cycle,omitempty" url:"accrued_this_cycle,omitempty"`
	StripeCustomerId     string   `protobuf:"bytes,4,opt,name=stripe_customer_id,json=stripeCustomerId,proto3" json:"stripe_customer_id,omitempty" db:"stripe_customer_id,omitempty" url:"stripe_customer_id,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty" db:"email,omitempty" url:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plan_Billing) GetMethod() BillingMethod {
	if m != nil {
		return m.Method
	}
	return BillingMethod_STRIPE
}

func (m *Plan_Billing) GetStripeCustomerId() string {
	if m != nil {
		return m.StripeCustomerId
	}
	return ""
}

type BillingMethod int32

const (
	BillingMethod_STRIPE      BillingMethod = 0
	BillingMethod_MANUAL      BillingMethod = 1
	BillingMethod_GCP         BillingMethod = 2
	BillingMethod_SKIP        BillingMethod = 3
	BillingMethod_MARKETPLACE BillingMethod = 4
	BillingMethod_NONE        BillingMethod = 5
)

var BillingMethod_name = map[int32]string{
	0: "STRIPE",
	1: "MANUAL",
	2: "GCP",
	3: "SKIP",
	4: "MARKETPLACE",
	5: "NONE",
}

var BillingMethod_value = map[string]int32{
	"STRIPE":      0,
	"MANUAL":      1,
	"GCP":         2,
	"SKIP":        3,
	"MARKETPLACE": 4,
	"NONE":        5,
}

func (x BillingMethod) String() string {
	return proto.EnumName(BillingMethod_name, int32(x))
}

type BillingInterval int32

const (
	BillingInterval_MONTHLY  BillingInterval = 0
	BillingInterval_ANNUALLY BillingInterval = 1
)

var BillingInterval_name = map[int32]string{
	0: "MONTHLY",
	1: "ANNUALLY",
}

var BillingInterval_value = map[string]int32{
	"MONTHLY":  0,
	"ANNUALLY": 1,
}

func (x BillingInterval) String() string {
	return proto.EnumName(BillingInterval_name, int32(x))
}

type SignupRequest struct {
	Organization         *Organization     `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	User                 *User             `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" db:"user,omitempty" url:"user,omitempty"`
	Credentials          *Credentials      `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty" db:"credentials,omitempty" url:"credentials,omitempty"`
	StripeToken          string            `protobuf:"bytes,4,opt,name=stripe_token,json=stripeToken,proto3" json:"stripe_token,omitempty" db:"stripe_token,omitempty" url:"stripe_token,omitempty"`
	Enterprise           bool              `protobuf:"varint,5,opt,name=enterprise,proto3" json:"enterprise,omitempty" db:"enterprise,omitempty" url:"enterprise,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,6,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Utm                  *Utm              `protobuf:"bytes,7,opt,name=utm,proto3" json:"utm,omitempty" db:"utm,omitempty" url:"utm,omitempty"`
	ExistingCreds        bool              `protobuf:"varint,8,opt,name=existing_creds,json=existingCreds,proto3" json:"existing_creds,omitempty" db:"existing_creds,omitempty" url:"existing_creds,omitempty"`
	Token                string            `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty" db:"token,omitempty" url:"token,omitempty"`
	MarketplaceCreds     *MarketplaceCreds `protobuf:"bytes,10,opt,name=marketplace_creds,json=marketplaceCreds,proto3" json:"marketplace_creds,omitempty" db:"marketplace_creds,omitempty" url:"marketplace_creds,omitempty"`
	CountryCode          string            `protobuf:"bytes,11,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty" db:"country_code,omitempty" url:"country_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SignupRequest) Reset()         { *m = SignupRequest{} }
func (m *SignupRequest) String() string { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()    {}

// TODO remove with CDMUM-789
type Credentials struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" db:"username,omitempty" url:"username,omitempty"`
	// to redact from json logs
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" redact:"-" db:"password,omitempty" url:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}

type Utm struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty" db:"source,omitempty" url:"source,omitempty"`
	Campaign             string   `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty" db:"campaign,omitempty" url:"campaign,omitempty"`
	Medium               string   `protobuf:"bytes,3,opt,name=medium,proto3" json:"medium,omitempty" db:"medium,omitempty" url:"medium,omitempty"`
	Partner              string   `protobuf:"bytes,4,opt,name=partner,proto3" json:"partner,omitempty" db:"partner,omitempty" url:"partner,omitempty"`
	Term                 string   `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty" db:"term,omitempty" url:"term,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty" db:"content,omitempty" url:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Utm) Reset()         { *m = Utm{} }
func (m *Utm) String() string { return proto.CompactTextString(m) }
func (*Utm) ProtoMessage()    {}

// TODO remove with CDMUM-789
type MarketplaceCreds struct {
	Token                string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" db:"token,omitempty" url:"token,omitempty"`
	Partner              MarketplacePartner `protobuf:"varint,2,opt,name=partner,proto3,enum=kafka.core.v1.MarketplacePartner" json:"partner,omitempty" db:"partner,omitempty" url:"partner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MarketplaceCreds) Reset()         { *m = MarketplaceCreds{} }
func (m *MarketplaceCreds) String() string { return proto.CompactTextString(m) }
func (*MarketplaceCreds) ProtoMessage()    {}

type SignupReply struct {
	User                 *User         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" db:"user,omitempty" url:"user,omitempty"`
	Organization         *Organization `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	Account              *Account      `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty" db:"account,omitempty" url:"account,omitempty"`
	Error                *Error        `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SignupReply) Reset()         { *m = SignupReply{} }
func (m *SignupReply) String() string { return proto.CompactTextString(m) }
func (*SignupReply) ProtoMessage()    {}

func (m *SignupReply) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

func (m *SignupReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type CreateEmailVerificationRequest struct {
	// Only username is required
	Credentials          *Credentials      `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty" db:"credentials,omitempty" url:"credentials,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,2,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateEmailVerificationRequest) Reset()         { *m = CreateEmailVerificationRequest{} }
func (m *CreateEmailVerificationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmailVerificationRequest) ProtoMessage()    {}

type Saml struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" db:"enabled,omitempty" url:"enabled,omitempty"`
	MetadataUrl          string   `protobuf:"bytes,2,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty" db:"metadata_url,omitempty" url:"metadata_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Saml) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.MetadataUrl)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Street1)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.Street2)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.Zip)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Plan_Billing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovOrg(uint64(m.Method))
	}
	if m.Interval != 0 {
		n += 1 + sovOrg(uint64(m.Interval))
	}
	if m.AccruedThisCycle != 0 {
		n += 1 + sovOrg(uint64(m.AccruedThisCycle))
	}
	l = len(m.StripeCustomerId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TaxAddress != nil {
		l = m.TaxAddress.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.ProductLevel != 0 {
		n += 1 + sovOrg(uint64(m.ProductLevel))
	}
	if m.TrialStart != nil {
		l = m.TrialStart.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.TrialEnd != nil {
		l = m.TrialEnd.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.PlanStart != nil {
		l = m.PlanStart.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.PlanEnd != nil {
		l = m.PlanEnd.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Product != nil {
		l = m.Product.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Billing != nil {
		l = m.Billing.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.ReferralCode)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.AcceptTos != nil {
		l = m.AcceptTos.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.AllowMultiTenant {
		n += 2
	}
	if m.AcceptTosPlatform != nil {
		l = m.AcceptTosPlatform.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Plan_Product) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReadThroughputMb != 0 {
		n += 9
	}
	if m.WriteThroughputMb != 0 {
		n += 9
	}
	if m.StorageGb != 0 {
		n += 1 + sovOrg(uint64(m.StorageGb))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Organization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovOrg(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Deactivated {
		n += 2
	}
	l = len(m.StripeCustomerId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Modified != nil {
		l = m.Modified.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.BillingEmail)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Plan != nil {
		l = m.Plan.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Saml != nil {
		l = m.Saml.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Sso != nil {
		l = m.Sso.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Marketplace != nil {
		l = m.Marketplace.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.ResourceId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.HasEntitlement {
		n += 2
	}
	if m.ShowBilling {
		n += 2
	}
	if m.AuditLog != nil {
		l = m.AuditLog.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.HasCommitment {
		n += 3
	}
	if m.MarketplaceSubscription != 0 {
		n += 2 + sovOrg(uint64(m.MarketplaceSubscription))
	}
	if m.DeactivatedAt != nil {
		l = m.DeactivatedAt.Size()
		n += 2 + l + sovOrg(uint64(l))
	}
	if m.SuspensionStatus != nil {
		l = m.SuspensionStatus.Size()
		n += 2 + l + sovOrg(uint64(l))
	}
	l = len(m.DisplayLabel)
	if l > 0 {
		n += 2 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *SuspensionStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Suspended != nil {
		l = m.Suspended.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovOrg(uint64(m.Status))
	}
	if m.EventType != 0 {
		n += 1 + sovOrg(uint64(m.EventType))
	}
	if m.ScheduledDeactivatedAt != nil {
		l = m.ScheduledDeactivatedAt.Size()
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovOrg(uint64(m.Version))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *AuditLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.AccountId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.ServiceAccountId != 0 {
		n += 1 + sovOrg(uint64(m.ServiceAccountId))
	}
	l = len(m.TopicName)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.ServiceAccountResourceId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Marketplace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Partner != 0 {
		n += 1 + sovOrg(uint64(m.Partner))
	}
	l = len(m.CustomerId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.CustomerState != 0 {
		n += 1 + sovOrg(uint64(m.CustomerState))
	}
	if m.ConsoleIntegrated {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Sso) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.Auth0ConnectionName)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	l = len(m.TenantId)
	if l > 0 {
		n += 1 + l + sovOrg(uint64(l))
	}
	if m.MultiTenant {
		n += 2
	}
	if m.Mode != 0 {
		n += 1 + sovOrg(uint64(m.Mode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TaxAddress != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.TaxAddress.Size()))
		n13, err := m.TaxAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.ProductLevel != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.ProductLevel))
	}
	if m.TrialStart != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.TrialStart.Size()))
		n14, err := m.TrialStart.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.TrialEnd != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.TrialEnd.Size()))
		n15, err := m.TrialEnd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.PlanStart != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.PlanStart.Size()))
		n16, err := m.PlanStart.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.PlanEnd != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.PlanEnd.Size()))
		n17, err := m.PlanEnd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.Product != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Product.Size()))
		n18, err := m.Product.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.Billing != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Billing.Size()))
		n19, err := m.Billing.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if len(m.ReferralCode) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.ReferralCode)))
		i += copy(dAtA[i:], m.ReferralCode)
	}
	if m.AcceptTos != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.AcceptTos.Size()))
		n20, err := m.AcceptTos.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.AllowMultiTenant {
		dAtA[i] = 0x60
		i++
		if m.AllowMultiTenant {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AcceptTosPlatform != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.AcceptTosPlatform.Size()))
		n21, err := m.AcceptTosPlatform.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n21
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Sso) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Enabled {
		dAtA[i] = 0x8
		i++
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Auth0ConnectionName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Auth0ConnectionName)))
		i += copy(dAtA[i:], m.Auth0ConnectionName)
	}
	if len(m.TenantId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.TenantId)))
		i += copy(dAtA[i:], m.TenantId)
	}
	if m.MultiTenant {
		dAtA[i] = 0x20
		i++
		if m.MultiTenant {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Mode != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Mode))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Marketplace) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Partner != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Partner))
	}
	if len(m.CustomerId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.CustomerId)))
		i += copy(dAtA[i:], m.CustomerId)
	}
	if m.CustomerState != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.CustomerState))
	}
	if m.ConsoleIntegrated {
		dAtA[i] = 0x20
		i++
		if m.ConsoleIntegrated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *AuditLog) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClusterId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.ClusterId)))
		i += copy(dAtA[i:], m.ClusterId)
	}
	if len(m.AccountId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.AccountId)))
		i += copy(dAtA[i:], m.AccountId)
	}
	if m.ServiceAccountId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.ServiceAccountId))
	}
	if len(m.TopicName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.TopicName)))
		i += copy(dAtA[i:], m.TopicName)
	}
	if len(m.ServiceAccountResourceId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.ServiceAccountResourceId)))
		i += copy(dAtA[i:], m.ServiceAccountResourceId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *SuspensionStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Suspended != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Suspended.Size()))
		n37, err := m.Suspended.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n37
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Status))
	}
	if m.EventType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.EventType))
	}
	if m.ScheduledDeactivatedAt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.ScheduledDeactivatedAt.Size()))
		n38, err := m.ScheduledDeactivatedAt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n38
	}
	if m.Version != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Version))
	}
	if len(m.ErrorMessage) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.ErrorMessage)))
		i += copy(dAtA[i:], m.ErrorMessage)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Saml) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Enabled {
		dAtA[i] = 0x8
		i++
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.MetadataUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.MetadataUrl)))
		i += copy(dAtA[i:], m.MetadataUrl)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Street1) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Street1)))
		i += copy(dAtA[i:], m.Street1)
	}
	if len(m.Street2) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Street2)))
		i += copy(dAtA[i:], m.Street2)
	}
	if len(m.City) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.City)))
		i += copy(dAtA[i:], m.City)
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.Country) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Country)))
		i += copy(dAtA[i:], m.Country)
	}
	if len(m.Zip) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Zip)))
		i += copy(dAtA[i:], m.Zip)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Plan_Billing) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Method != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Method))
	}
	if m.Interval != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Interval))
	}
	if m.AccruedThisCycle != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.AccruedThisCycle))
	}
	if len(m.StripeCustomerId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.StripeCustomerId)))
		i += copy(dAtA[i:], m.StripeCustomerId)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Plan_Product) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReadThroughputMb != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ReadThroughputMb))))
		i += 8
	}
	if m.WriteThroughputMb != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.WriteThroughputMb))))
		i += 8
	}
	if m.StorageGb != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.StorageGb))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Organization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Deactivated {
		dAtA[i] = 0x18
		i++
		if m.Deactivated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.StripeCustomerId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.StripeCustomerId)))
		i += copy(dAtA[i:], m.StripeCustomerId)
	}
	if m.Created != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Created.Size()))
		n28, err := m.Created.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n28
	}
	if m.Modified != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Modified.Size()))
		n29, err := m.Modified.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n29
	}
	if len(m.BillingEmail) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.BillingEmail)))
		i += copy(dAtA[i:], m.BillingEmail)
	}
	if m.Plan != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Plan.Size()))
		n30, err := m.Plan.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n30
	}
	if m.Saml != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Saml.Size()))
		n31, err := m.Saml.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n31
	}
	if m.Sso != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Sso.Size()))
		n32, err := m.Sso.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n32
	}
	if m.Marketplace != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.Marketplace.Size()))
		n33, err := m.Marketplace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n33
	}
	if len(m.ResourceId) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.ResourceId)))
		i += copy(dAtA[i:], m.ResourceId)
	}
	if m.HasEntitlement {
		dAtA[i] = 0x68
		i++
		if m.HasEntitlement {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ShowBilling {
		dAtA[i] = 0x70
		i++
		if m.ShowBilling {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AuditLog != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.AuditLog.Size()))
		n34, err := m.AuditLog.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n34
	}
	if m.HasCommitment {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		if m.HasCommitment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MarketplaceSubscription != 0 {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.MarketplaceSubscription))
	}
	if m.DeactivatedAt != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.DeactivatedAt.Size()))
		n35, err := m.DeactivatedAt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n35
	}
	if m.SuspensionStatus != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintOrg(dAtA, i, uint64(m.SuspensionStatus.Size()))
		n36, err := m.SuspensionStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n36
	}
	if len(m.DisplayLabel) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintOrg(dAtA, i, uint64(len(m.DisplayLabel)))
		i += copy(dAtA[i:], m.DisplayLabel)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaxAddress == nil {
				m.TaxAddress = &Address{}
			}
			if err := m.TaxAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductLevel", wireType)
			}
			m.ProductLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProductLevel |= ProductLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrialStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TrialStart == nil {
				m.TrialStart = &types.Timestamp{}
			}
			if err := m.TrialStart.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrialEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TrialEnd == nil {
				m.TrialEnd = &types.Timestamp{}
			}
			if err := m.TrialEnd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlanStart == nil {
				m.PlanStart = &types.Timestamp{}
			}
			if err := m.PlanStart.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlanEnd == nil {
				m.PlanEnd = &types.Timestamp{}
			}
			if err := m.PlanEnd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Product", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Product == nil {
				m.Product = &Plan_Product{}
			}
			if err := m.Product.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Billing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Billing == nil {
				m.Billing = &Plan_Billing{}
			}
			if err := m.Billing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferralCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReferralCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptTos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcceptTos == nil {
				m.AcceptTos = &types.BoolValue{}
			}
			if err := m.AcceptTos.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowMultiTenant", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowMultiTenant = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptTosPlatform", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcceptTosPlatform == nil {
				m.AcceptTosPlatform = &types.BoolValue{}
			}
			if err := m.AcceptTosPlatform.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Organization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Organization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Organization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deactivated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deactivated = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StripeCustomerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StripeCustomerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &types.Timestamp{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modified", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Modified == nil {
				m.Modified = &types.Timestamp{}
			}
			if err := m.Modified.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillingEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BillingEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Plan == nil {
				m.Plan = &Plan{}
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Saml", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Saml == nil {
				m.Saml = &Saml{}
			}
			if err := m.Saml.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sso", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sso == nil {
				m.Sso = &Sso{}
			}
			if err := m.Sso.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Marketplace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Marketplace == nil {
				m.Marketplace = &Marketplace{}
			}
			if err := m.Marketplace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasEntitlement", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasEntitlement = bool(v != 0)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShowBilling", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ShowBilling = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuditLog", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuditLog == nil {
				m.AuditLog = &AuditLog{}
			}
			if err := m.AuditLog.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasCommitment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasCommitment = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketplaceSubscription", wireType)
			}
			m.MarketplaceSubscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketplaceSubscription |= MarketplaceSubscription(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeactivatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeactivatedAt == nil {
				m.DeactivatedAt = &types.Timestamp{}
			}
			if err := m.DeactivatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuspensionStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuspensionStatus == nil {
				m.SuspensionStatus = &SuspensionStatus{}
			}
			if err := m.SuspensionStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Street1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Street1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Street2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Street2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Plan_Product) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Product: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Product: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadThroughputMb", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.ReadThroughputMb = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriteThroughputMb", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.WriteThroughputMb = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageGb", wireType)
			}
			m.StorageGb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageGb |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Saml) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Saml: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Saml: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Plan_Billing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Billing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Billing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= BillingMethod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interval |= BillingInterval(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccruedThisCycle", wireType)
			}
			m.AccruedThisCycle = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccruedThisCycle |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StripeCustomerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StripeCustomerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Sso) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Sso: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sso: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auth0ConnectionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auth0ConnectionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiTenant", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MultiTenant = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= SsoMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Marketplace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Marketplace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Marketplace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partner", wireType)
			}
			m.Partner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Partner |= MarketplacePartner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustomerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerState", wireType)
			}
			m.CustomerState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CustomerState |= MarketplaceCustomerState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsoleIntegrated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ConsoleIntegrated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuditLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuditLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuditLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountId", wireType)
			}
			m.ServiceAccountId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceAccountId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TopicName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountResourceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAccountResourceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SuspensionStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SuspensionStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuspensionStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspended", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Suspended == nil {
				m.Suspended = &types.Timestamp{}
			}
			if err := m.Suspended.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SuspensionStatusType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= SuspensionEventType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledDeactivatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScheduledDeactivatedAt == nil {
				m.ScheduledDeactivatedAt = &types.Timestamp{}
			}
			if err := m.ScheduledDeactivatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func encodeVarintOrg(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func sovOrg(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func skipOrg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrg
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOrg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOrg
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOrg(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOrg
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOrg = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrg   = fmt.Errorf("proto: integer overflow")
)
