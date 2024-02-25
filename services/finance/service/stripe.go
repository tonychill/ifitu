package service

import (
	"context"
	"log"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/tonychill/ifitu/apis/pb/go/global"
)

type paymentIntent struct {
	clientSecret string
}

func (s *ServiceImpl) setupPayment(ctx context.Context, conf *global.Confirmation) (paymentIntent, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// customers, err := s.repo.GetCustomers(ctx, &global.Query{
	// 	Terms: []*global.Term{
	// 		{
	// 			Key:   "id",
	// 			Value: "guest id from context",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Error().Err(err).Msg("error getting guest from finance repository")
	// 	return paymentIntent{}, err
	// }

	// fmt.Println(guests)

	// get the data needed to create a payment intnet

	return paymentIntent{}, nil
}

type checkoutSession struct {
	Url string
}

func (s *ServiceImpl) createCheckoutSession(ctx context.Context) (checkoutSession, error) {
	domain := "https://ifitu.fly.dev"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1OnYN9IewGpSuDNGLLlk3vIT"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:   stripe.String(domain + "/success.html"),
		CancelURL:    stripe.String(domain + "/cancel.html"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}

	sesh, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
	}

	return checkoutSession{
		Url: sesh.URL,
	}, nil

}
