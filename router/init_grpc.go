package router

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	// conSvc "github.com/tonychill/ifitu/apis/pb/go/coordinator_service"
	// contSvc "github.com/tonychill/ifitu/apis/pb/go/content_service"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/apis/pb/go/health_check"
	// idSvc "github.com/tonychill/ifitu/apis/pb/go/identity_service"


	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	// "github.com/ilanlyfe/lib/traceutil"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/sdk/trace"
)

// Registers grpc service implementations with the router, starts a listener on the
// specified grpc port, and starts the grpc server.
func (r *routerImpl) initGrpcServer(ctx context.Context, g *errgroup.Group) error {

	ep := &entryPoint{
		ready: true,
	}

	grpcServer := createGrpcServer()
	r.grpcSvrShutdown = func() {
		log.Info().Msg("Shutting down GRPC servers...")
		grpcServer.GracefulStop()
	}

	g.Go(func() error {
		// conSvc.RegisterConciergeServiceServer(grpcServer, r.conSvc)
		health_check.RegisterHealthServer(grpcServer, ep)

		// contSvc.RegisterContentServiceServer(grpcServer, r.contImpl)
		// idSvc.RegisterIdentityServiceServer(grpcServer, r.idImpl)
		finSvc.RegisterFinanceServiceServer(grpcServer, r.finImpl)

		lis, err := net.Listen("tcp", r.GrpcServerPort)
		if err != nil {
			log.Error().Msgf("failed to listen on port %s during grpc sever initialization: %v",
				r.GrpcServerPort, err)
			return err
		}

		log.Info().Msgf("GRPC server listening on port %s", r.GrpcServerPort)
		if err = grpcServer.Serve(lis); err != nil {
			log.Error().Msg(fmt.Sprintf("failed to serve: %v", err))
			return err
		}
		return nil
	})
	return nil
}

func createGrpcServer(interceptors ...grpc.UnaryServerInterceptor) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors...,
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(),
		),
	)
	return s
}

// TODO: implement a channel that will be used to communicate with the
// the services so that they are able to let the entrypoint know when
// they are all up and ready to serve requests.
type entryPoint struct {
	ready bool
	// live  bool TODO: implement
	health_check.HealthServer
}

func (e *entryPoint) Check(_ context.Context, r *health_check.HealthCheckRequest) (*health_check.HealthCheckResponse, error) {
	response := &health_check.HealthCheckResponse{
		Status: health_check.HealthCheckResponse_SERVING_STATUS_SERVING,
	}
	// a liveness check determines that the service is up, but not
	// necessarily 'ready'
	if r.Service == "liveness" {
		switch e.ready {
		case true:
			response.Status = health_check.HealthCheckResponse_SERVING_STATUS_SERVING
		default:
			log.Warn().Msg("services are not in a ready state")
			response.Status = health_check.HealthCheckResponse_SERVING_STATUS_NOT_SERVING
		}
	} else { // if we get here kubernetes asked for a 'readiness' check
		if e.ready {
			response.Status = health_check.HealthCheckResponse_SERVING_STATUS_NOT_SERVING
		} else {
			log.Warn().Msg("services are not in a ready state")
			response.Status = health_check.HealthCheckResponse_SERVING_STATUS_SERVING
		}
	}

	return response, nil
}

func (s *entryPoint) Watch(_ *health_check.HealthCheckRequest, _ health_check.Health_WatchServer) error {
	return nil
}
