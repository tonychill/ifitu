package router

import (
	"context"
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/tonychill/ifitu/lib/grpcutil"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func InitializeRouter(ctx context.Context, g *errgroup.Group,
	impl ServiceImplementations) (Router, error) {
	if impl.CoordinatorService == nil {
		return nil, fmt.Errorf("a concierge service was not provided when attempting to initialize the router")
	}

	r := &routerImpl{
		corImpl: impl.CoordinatorService,
		// The follwoing services should not be called directly
		// contImpl: impl.ContentService,
		finImpl: impl.FinanceService,
		// idImpl:   impl.IdentityService,

	}

	if err := envconfig.Process("", r); err != nil {
		return nil, err
	}

	if err := r.initHttpServer(ctx, g); err != nil {
		return nil, fmt.Errorf("could not initiate http router: %s", err.Error())
	}

	if err := r.initGrpcServer(ctx, g); err != nil {
		return nil, fmt.Errorf("could not initiate grpc server: %s", err.Error())
	}

	r, err := r.connectToServices(ctx, g)
	if err != nil {
		log.Error().Msgf("failed to connect interfaces: %s", err.Error())
		return nil, err
	}

	return r, nil
}

func (r *routerImpl) connectToServices(ctx context.Context, g *errgroup.Group) (*routerImpl, error) {

	// Pull in the agent's configuration from the environment

	// TODO:
	// Based on the lead agent's placement of various services, on startup the router will
	// decide wether or not it is to handle implementation requests interprocess or if it
	// needs to connect to another instance of an agent over the wire. However, this does
	// not mean that the router will not be able to handle requests from other agents. Also,
	// the router should be able to dynamically change its implementation based on the
	// current state of the system. For example, if the router is currently handling
	// requests for a service but the local resources hits a threshold, the router should
	// be able to hand off the implementation to another agent. This will require the router
	// to be able to communicate with other agents and be able to determine the current
	// state of the system.
	// go func() {
	// 	if err := r.connectToServices(ctx); err != nil {
	// 		log.Fatal().Msgf("failed to connect to services: %v", err)
	// 	}
	// }()

	conn, err := grpcutil.NewGrpcClientConn(grpcutil.ConnectToGrpcServiceParams{
		// TODO: pull the following values from the service config
		Name:    "concierge-service",
		Address: "localhost:50001",
	})
	if err != nil {
		log.Err(err).Msg("failed to setup connection to")
	}

	if err := r.corImpl.ConnectClients(ctx, conn); err != nil {
		log.Error().Err(err).Msg("failed to connect coordinator to services during router initialization")
		return nil, fmt.Errorf("failed to connect to coordinator service: %w", err)
	}

	return r, nil
}

// TODO: in work
func (s *routerImpl) _connectToServices(ctx context.Context) error {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		grpcutil.ConnectToGrpcService(
			grpcutil.ConnectToGrpcServiceParams{
				// Logger:  any,
				Name: "operations-service",
				// Quit:    s.shutdownCh,
				// Address: s.OperationsServiceAddress,
				// OnConnect: func(conn grpc.ClientConnInterface) {
				// 	s.opsSvc = opsPb.NewOperationsServiceClient(conn)
				// },
			})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		grpcutil.ConnectToGrpcService(
			grpcutil.ConnectToGrpcServiceParams{
				// Logger:  any,
				Name: "journeys-service",
				// Quit:    s.shutdownCh,
				// Address: s.JourneyServiceAddress,
				// OnConnect: func(conn grpc.ClientConnInterface) {
				// 	s.journeyClient = jnySvc.NewJourneysServiceClient(conn)
				// 	s.conciergeClient = conSvc.NewConciergeServiceClient(conn)
				// },
			})
	}(&wg)

	wg.Wait()

	log.Info().Msg("connected to services")

	return nil

}
