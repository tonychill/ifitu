package grpcutil

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tonychill/ifitu/apis/pb/go/health_check"
)

// GRPCServer describes the required methods a gRPC server must implement in order to be managed
type GRPCServer interface {
	health_check.HealthServer
	Shutdown(context.Context)
}

// GRPCService wraps a GRPCServer in order to provide consistent functionality (e.g. running, logging, and shutting down)
type GRPCService struct {
	// Logger           log.Logger
	Listen           string
	UIListen         string
	Server           GRPCServer
	ServiceDesc      *grpc.ServiceDesc
	ShutdownDuration time.Duration
}

// RunGrpcServer accepts a *grpc.Server without configuring it. It provides informative logs, and
// handles graceful shutdowns.
func (s GRPCService) RunGrpcServer(grpcServer *grpc.Server) (errorCode int) {
	// s.Logger.Info().Msg("starting gRPC server...")
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", s.Listen)
	if err != nil {
		// s.Logger.Err(err).Msg("failed to listen")
		return 1
	}

	var adminShutdown func(context.Context) []error
	var uiErr error
	// var uiServer *http.Server

	if s.UIListen == "" {
		// s.Logger.Info().Msg("grpc ui not enabled. skipping")
	} else {
		go func() {
			_, adminShutdown, uiErr = newGRPCUIServer(
				context.Background(),
				s.UIListen,
				s.Listen,
			)
			if uiErr != nil {
				// s.Logger.Err(uiErr).Msg("failed to start gRPC UI server")
				return
			}
			// s.Logger.Info().Msgf("gRPC UI server running at %s", s.UIListen)
			// s.Logger.Err(uiServer.ListenAndServe()).Msg("gRPC UI server stopped")
		}()
	}

	go s.handleInterrupt(grpcServer, adminShutdown)

	// s.Logger.Info().Msgf("gRPC server running at %s", s.Listen)
	if err := grpcServer.Serve(lis); err != nil {
		// s.Logger.Err(err).Msg("failed to serve")
		errorCode = 1 // set errorCode to 1 since we're shutting down due to a fatal error (not an interrupt signal)
	}

	s.shutdown()
	// s.Logger.Warn().Msg("grpc server stopped")
	return errorCode // errorCode will still be 0 if we're shutting down from an interrupt signal
}

// Run attempts to create a gRPC server in a consistent way and then run it. It provides informative
// logs, and handles graceful shutdowns.
func (s GRPCService) Run(serverOptions ...grpc.ServerOption) (errorCode int) {
	opts := append([]grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 64),
	}, serverOptions...)
	grpcServer := grpc.NewServer(opts...)
	grpcServer.RegisterService(s.ServiceDesc, s.Server)

	health_check.RegisterHealthServer(grpcServer, s.Server)
	return s.RunGrpcServer(grpcServer)
}

func (s *GRPCService) shutdown() {
	// s.Logger.Info().Msg("initiating shutdown")

	duration := 5 * time.Second
	if s.ShutdownDuration != 0 {
		duration = s.ShutdownDuration
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	shutdownComplete := make(chan struct{})
	go func() {
		s.Server.Shutdown(ctx)
		close(shutdownComplete)
	}()

	select {
	case <-ctx.Done():
		// s.Logger.Warn().Msg("aborting shutdown")
		return
	case <-shutdownComplete:
		return
	}
}

func (s *GRPCService) handleInterrupt(
	grpcServer *grpc.Server,
	adminShutdown func(context.Context) []error,
) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	<-interrupt

	// s.Logger.Warn().Msg("received shutdown signal")

	if adminShutdown != nil {
		duration := 5 * time.Second
		if s.ShutdownDuration != 0 {
			duration = s.ShutdownDuration
		}
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()
		for range adminShutdown(ctx) {
			// s.Logger.Err(err).Msg("grpc ui shutdown error")
		}
	}

	grpcServer.GracefulStop()
}

// func createGrpcServer(interceptors ...grpc.UnaryServerInterceptor) *grpc.Server {
// 	return grpc.NewServer(
// 		grpc.UnaryInterceptor(
// 			grpc_middleware.ChainUnaryServer(
// 				interceptors...,
// 			),
// 		),
// 		grpc.StreamInterceptor(
// 			grpc_middleware.ChainStreamServer(),
// 		),
// 	)

// }
