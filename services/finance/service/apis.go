package service

import (
	"context"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/setupintent"
	"github.com/tonychill/ifitu/apis/pb/go/finance"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
)

func (s *ServiceImpl) GetPayments(ctx context.Context, req *finSvc.GetPaymentsRequest) (*finSvc.GetPaymentsResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) AddPaymentMethod(ctx context.Context, req *finSvc.AddPaymentMethodRequest) (*finSvc.AddPaymentMethodResponse, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.SetupIntentParams{
		Usage: stripe.String(string(stripe.SetupIntentUsageOffSession)),
	}
	result, err := setupintent.New(params)
	if err != nil {
		return nil, err
	}

	// if err := s.repo.AddPaymentMethod(ctx, req); err != nil {
	// 	log.Error().Err(err).Msgf("failed to add payment method for guest id %s",
	// 		req.PaymentMethod.GuestId)
	// 	return nil, err
	// }
	return &finSvc.AddPaymentMethodResponse{
		PaymentMethod: &finance.PaymentMethod{
			CustomerId: result.ClientSecret,
		},
	}, nil
}

func (s *ServiceImpl) GetPaymentMethods(ctx context.Context, req *finSvc.GetPaymentMethodsRequest) (*finSvc.GetPaymentMethodsResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) RemovePaymentMethod(ctx context.Context, req *finSvc.RemovePaymentMethodRequest) (*finSvc.RemovePaymentMethodResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) SaveAuthorization(ctx context.Context, req *finSvc.SaveAuthorizationRequest) (*finSvc.SaveAuthorizationResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) CaptureFunds(ctx context.Context, req *finSvc.CaptureFundsRequest) (*finSvc.CaptureFundsResponse, error) {
	panic("implement me please")
}

func (s *ServiceImpl) Shutdown(ctx context.Context) error {
	s.shutdown = true
	s.repo.Shutdown(ctx)
	return nil
}
