package ccloud

import (
	"context"
)

// Auth allows authenticating in Confluent Cloud
type Auth interface {
	Login(context.Context, *AuthenticateRequest) (*AuthenticateReply, error)
	User(context.Context) (*GetMeReply, error)
}

// Signup service allows managing signups in Confluent Cloud
type Signup interface {
	Create(context.Context, *SignupRequest) (*SignupReply, error)
	SendVerificationEmail(context.Context, *User) error
}

// Billing service allows getting billing information for an org in Confluent Cloud
type Billing interface {
	GetPriceTable(ctx context.Context, org *Organization, product string) (*PriceTable, error)
	GetPaymentInfo(ctx context.Context, org *Organization) (*Card, error)
	UpdatePaymentInfo(ctx context.Context, org *Organization, stripeToken string) error
	ClaimPromoCode(ctx context.Context, org *Organization, code string) (*PromoCodeClaim, error)
	GetClaimedPromoCodes(ctx context.Context, org *Organization, excludeExpired bool) ([]*PromoCodeClaim, error)
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
