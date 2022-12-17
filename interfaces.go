package ccloud

import (
	"context"
)

// Account service allows managing accounts in Confluent Cloud
type AccountInterface interface {
	Create(context.Context, *Account) (*Account, error)
	Get(context.Context, *Account) (*Account, error)
	List(context.Context, *Account) ([]*Account, error)
}

// Auth allows authenticating in Confluent Cloud
type Auth interface {
	Login(context.Context, *AuthenticateRequest) (*AuthenticateReply, error)
	User(context.Context) (*GetMeReply, error)
}

// Billing service allows getting billing information for an org in Confluent Cloud
type Billing interface {
	GetPriceTable(ctx context.Context, org *Organization, product string) (*PriceTable, error)
	GetPaymentInfo(ctx context.Context, org *Organization) (*Card, error)
	UpdatePaymentInfo(ctx context.Context, org *Organization, stripeToken string) error
	ClaimPromoCode(ctx context.Context, org *Organization, code string) (*PromoCodeClaim, error)
	GetClaimedPromoCodes(ctx context.Context, org *Organization, excludeExpired bool) ([]*PromoCodeClaim, error)
}

// Environment metadata service allows getting information about available cloud regions data
type EnvironmentMetadata interface {
	Get(context.Context) ([]*CloudMetadata, error)
}

// External Identity services allow managing external identities for Bring-Your-Own-Key in Confluent Cloud.
type ExternalIdentity interface {
	CreateExternalIdentity(ctx context.Context, cloud, accountID string) (externalIdentityName string, err error)
}

type Growth interface {
	GetFreeTrialInfo(context.Context, int32) ([]*GrowthPromoCodeClaim, error)
}

// Schema Registry service allows managing SR clusters in Confluent Cloud
type SchemaRegistry interface {
	CreateSchemaRegistryCluster(context.Context, *SchemaRegistryClusterConfig) (*SchemaRegistryCluster, error)
	GetSchemaRegistryClusters(context.Context, *SchemaRegistryCluster) ([]*SchemaRegistryCluster, error)
	GetSchemaRegistryCluster(context.Context, *SchemaRegistryCluster) (*SchemaRegistryCluster, error)
	UpdateSchemaRegistryCluster(context.Context, *SchemaRegistryCluster) (*SchemaRegistryCluster, error)
	DeleteSchemaRegistryCluster(context.Context, *SchemaRegistryCluster) error
}

// Signup service allows managing signups in Confluent Cloud
type Signup interface {
	Create(context.Context, *SignupRequest) (*SignupReply, error)
	SendVerificationEmail(context.Context, *User) error
}

// User service allows managing users in Confluent Cloud
type UserInterface interface {
	List(context.Context) ([]*User, error)
	GetServiceAccounts(context.Context) ([]*User, error)
	GetServiceAccount(context.Context, int32) (*User, error)
	LoginRealm(context.Context, *GetLoginRealmRequest) (*GetLoginRealmReply, error)
}

// Logger provides an interface that will be used for all logging in this client. User provided
// logging implementations must conform to this interface. Popular loggers like zap and logrus
// already implement this interface.
type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}
