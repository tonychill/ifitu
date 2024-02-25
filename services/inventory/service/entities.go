package service

import (
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/services/finance/repo"
)

var (
	_ = finSvc.FinanceServiceServer(&ServiceImpl{})
)

type ServiceImpl struct {
	shutdownCh chan struct{}
	shutdown   bool
	repo       repo.Repository
	finSvc.UnimplementedFinanceServiceServer
}
