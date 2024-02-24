package service

import (
	"context"
	"os"

	"github.com/stripe/stripe-go/v76"
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
