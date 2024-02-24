package service

import (
	conSvr "github.com/tonychill/ifitu/apis/pb/go/coordinator_service"
	finClient "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/services/coordinator/repo"
)

type Server interface {

	// Send(*conSvc.FlowStatus) error
	// grpc.ServerStream
}

var (
	_ = conSvr.CoordinatorServiceServer(&ServiceImpl{})
)

type ServiceImpl struct {
	conSvr.UnimplementedCoordinatorServiceServer
	ready      bool
	shutdownCh chan struct{}
	repo       repo.Repository
	// contentClient cntClient.ContentServiceClient
	// idClient      idClient.IdentityServiceClient
	financeClient finClient.FinanceServiceClient
}
