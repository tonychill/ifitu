package service

import (
	"context"
	"fmt"

	// contSvc "github.com/tonychill/ifitu/apis/pb/go/content_service"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	// idSvc "github.com/tonychill/ifitu/apis/pb/go/identity_service"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/services/coordinator/repo"
	"google.golang.org/grpc"
)

func InitializeConciergeService(ctx context.Context) (*ServiceImpl, error) {
	repo, err := repo.InitiateRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("the repository failed to initialize during coordinator service startup: %w", err)
	}

	s := &ServiceImpl{
		repo:       repo,
		shutdownCh: make(chan struct{}),
	}

	if err := envconfig.Process("", s); err != nil {
		log.Fatal().Msgf("error when processing config from env: %+v",
			err)
	}

	// TODO: handle cases where the coordinator may no be able to connect to a
	// one or more services
	s.ready = true
	return s, nil
}

// Connects the coordinator service to the other downstream services. This method should
// only be called by the router during startup.
func (s *ServiceImpl) ConnectClients(ctx context.Context, conn *grpc.ClientConn) error {
	// s.contentClient = contSvc.NewContentServiceClient(conn)
	// s.idClient = idSvc.NewIdentityServiceClient(conn)
	s.financeClient = finSvc.NewFinanceServiceClient(conn)
	return nil
}
