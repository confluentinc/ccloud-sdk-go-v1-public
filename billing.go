package ccloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	priceTableEndpoint                 = "/api/organizations/%d/price_table?product=%v"
	paymentInfoEndpoint                = "/api/organizations/%d/payment_info"
	updateDefaultPaymentMethodEndpoint = "/api/organizations/%d/update_default_payment_method"
	billingProfileEndpoint             = "/api/organizations/%d/billing_profile"
	taxInfoEndpoint                    = "/api/organizations/%d/tax_info"
	promoCodeEndpoint                  = "/api/organizations/%d/promo_code_claims"
)

type BillingService struct {
	client *http.Client
	sling  *sling.Sling
}

func NewBillingService(client *Client) *BillingService {
	return &BillingService{
		client: client.HttpClient,
		sling:  client.sling,
	}
}

func (o *BillingService) GetPriceTable(_ context.Context, org *Organization, product string) (*PriceTable, error) {
	url := fmt.Sprintf(priceTableEndpoint, org.Id, product)
	res := new(GetPriceTableReply)

	_, err := o.sling.New().Get(url).Receive(res, res)
	if err := ReplyErr(res, err); err != nil {
		return nil, WrapErr(err, "error getting price table")
	}

	return res.PriceTable, nil
}

func (o *BillingService) GetPaymentInfo(_ context.Context, org *Organization) (*Card, error) {
	url := fmt.Sprintf(paymentInfoEndpoint, org.Id)
	res := new(GetPaymentInfoReply)

	_, err := o.sling.New().Get(url).Receive(res, res)
	return res.Card, WrapErr(ReplyErr(res, err), "failed to get payment info")
}

func (o *BillingService) UpdatePaymentInfo(_ context.Context, org *Organization, stripeToken string) error {
	url := fmt.Sprintf(paymentInfoEndpoint, org.Id)
	req := &UpdatePaymentInfoRequest{StripeToken: stripeToken}
	res := &UpdatePaymentInfoReply{}

	_, err := o.sling.New().Post(url).BodyProvider(Request(req)).Receive(res, res)
	return WrapErr(ReplyErr(res, err), "failed to update payment info")
}

func (o *BillingService) ClaimPromoCode(_ context.Context, org *Organization, code string) (*PromoCodeClaim, error) {
	url := fmt.Sprintf(promoCodeEndpoint, org.Id)
	req := &ClaimPromoCodeRequest{Code: code}
	res := &ClaimPromoCodeReply{}

	_, err := o.sling.New().Post(url).BodyProvider(Request(req)).Receive(res, res)
	return res.Claim, WrapErr(ReplyErr(res, err), "failed to claim promo code")
}

func (o *BillingService) GetClaimedPromoCodes(_ context.Context, org *Organization, excludeExpired bool) ([]*PromoCodeClaim, error) {
	url := fmt.Sprintf(promoCodeEndpoint, org.Id)
	req := &GetPromoCodeClaimsRequest{ExcludeExpired: excludeExpired}
	res := &GetPromoCodeClaimsReply{}

	_, err := o.sling.New().Get(url).BodyProvider(Request(req)).Receive(res, res)
	return res.Claims, WrapErr(ReplyErr(res, err), "failed to get claimed promo codes")
}
