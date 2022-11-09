package ccloud

import (
	"context"

	orgv1 "github.com/confluentinc/cc-structs/kafka/org/v1"
)

// Billing service allows getting billing information for an org in Confluent Cloud
type Billing interface {
	GetPriceTable(ctx context.Context, org *orgv1.Organization, product string) (*PriceTable, error)
	GetPaymentInfo(ctx context.Context, org *orgv1.Organization) (*Card, error)
	UpdatePaymentInfo(ctx context.Context, org *orgv1.Organization, stripeToken string) error
	ClaimPromoCode(ctx context.Context, org *orgv1.Organization, code string) (*PromoCodeClaim, error)
	GetClaimedPromoCodes(ctx context.Context, org *orgv1.Organization, excludeExpired bool) ([]*PromoCodeClaim, error)
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
