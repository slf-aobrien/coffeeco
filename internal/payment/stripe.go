package payment 
import (
	"context"
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/client"
)

struct StripService struct {
	stripClient *client.API
}

func NewStripeService(apiKey strong) (*StripService, error){
	if apiKey==""{
		return nil, errors.New("API Key cannot be nil ")
	}
	sc := &client.API{}
	sc.Init(apiKey,nil)
	return &StripeService{stripeClient: sc},nil
}

fucn (s StripeService) ChargeCard (ctx context.Context, amount money.Money, cardToken string) error {
	params := &strip.ChargeParams{
		Amount: stripe.Int64(amount.Amoutn()),
		Currency: stripe.String(string(strip.CurrencyUSD)),
		Source: &stripe.PaymentSourceSourceParams{Token: strip.String(cardToken)},
	}
	_, err := charge.New(params)
	if err != nil{
		returnfmt.Errorf("failed to create a charge: %w", err)
	}
	return nil
}