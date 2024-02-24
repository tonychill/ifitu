package repo

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	finPb "github.com/tonychill/ifitu/apis/pb/go/finance"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/apis/pb/go/global"
)

type mockRepoImpl struct {
	sync.RWMutex
	index map[string]*global.Rate
}

func newMockRepository() (Repository, error) {
	return &mockRepoImpl{
		index: make(map[string]*global.Rate),
	}, nil
}

func (r *mockRepoImpl) AddPaymentMethod(ctx context.Context, req *finSvc.AddPaymentMethodRequest) error {
	panic("implement me")
}
func (r *mockRepoImpl) GetPaymentMethods(ctx context.Context, req *finSvc.GetPaymentMethodsRequest) ([]*finPb.PaymentMethod, error) {
	panic("implement me")
}
func (r *mockRepoImpl) RemovePaymentMethod(ctx context.Context, req *finSvc.RemovePaymentMethodRequest) error {
	panic("implement me")
}
func (r *mockRepoImpl) GetRates(ctx context.Context, req *global.Query) ([]*global.Rate, error) {
	panic("not implemented")
}

func (r *mockRepoImpl) CreateRate(ctx context.Context, req *global.Rate) (string, error) {
	log.Debug().Msgf("creating mock resource: %+v\n", req)
	// id, err := lib.NewULID("")
	// if err != nil {
	// 	return "", err
	// }
	// r.index[id] = nil
	// return id, nil
	panic("not implemented")
}

func (r *mockRepoImpl) Shutdown(ctx context.Context) error {
	return nil
}
