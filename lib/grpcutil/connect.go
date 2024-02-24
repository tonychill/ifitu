package grpcutil

import (
	"context"
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/rs/zerolog/log"
)

const MaxMsgSize = 67108864 //64MB = 67108864 bytes

var baseUnaryClientInterceptors = []grpc.UnaryClientInterceptor{
	// clientmiddleware.ContextClientInterceptor,
}

var baseStreamClientInterceptors = []grpc.StreamClientInterceptor{
	// clientmiddleware.ContextClientStreamingInterceptor,
}

type ConnectToGrpcServiceParams struct {
	Address                  string
	UnaryClientInterceptors  []grpc.UnaryClientInterceptor
	StreamClientInterceptors []grpc.StreamClientInterceptor
	// Logger                   log.Logger
	Name      string
	Quit      <-chan struct{}
	OnConnect func(conn grpc.ClientConnInterface)
}

// ConnectToGrpcService attempts to connect to a gRPC address and return a client connection. This method will wait
// indefinitely for a connection or a quit signal.
func ConnectToGrpcService(params ConnectToGrpcServiceParams) {
	if params.OnConnect == nil {
		log.Error().Msg("OnConnect parameter is required")
		return
	}
	log.Info().Msgf("ConnectToGrpcService: params: %+v", params)
	dialOpts := newDialOptions(params)
	dialOpts = append(dialOpts, grpc.WithBlock())

	var clientConn *grpc.ClientConn
	var err error

Loop:
	for {
		select {
		case <-params.Quit:
			log.Info().Msgf("quitting attempts to connect to %s", params.Name)
			break Loop
		default:
			log.Info().Msgf("connecting to %s at %s", params.Name, params.Address)
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
			clientConn, err = grpc.DialContext(ctx, params.Address, dialOpts...)
			cancel()
			if err != nil {
				log.Err(err).Msgf("could not connect to %s at %s. retrying", params.Name, params.Address)
				continue
			}
			log.Info().Msgf("successfully connected to %s at %s", params.Name, params.Address)
			params.OnConnect(clientConn)
			break Loop
		}
	}
}

func newDialOptions(params ConnectToGrpcServiceParams) []grpc.DialOption {
	transportCred := grpc.WithTransportCredentials(insecure.NewCredentials())
	if strings.HasSuffix(params.Address, "443") {
		transportCred = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	}

	unaryClientInterceptors := buildUnaryClientInterceptors(params)
	streamClientInterceptors := buildStreamClientInterceptors(params)

	return []grpc.DialOption{
		transportCred,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(MaxMsgSize),
			grpc.MaxCallSendMsgSize(MaxMsgSize),
		),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             30 * time.Second,
				PermitWithoutStream: true,
			}),
		grpc.WithChainUnaryInterceptor(unaryClientInterceptors...),
		grpc.WithChainStreamInterceptor(streamClientInterceptors...),
	}
}

// buildStreamClientInterceptors returns the default list of interceptors for a stream client connection.
// Additional interceptors can be added to this list by passing in params.StreamClientInterceptors.
func buildStreamClientInterceptors(params ConnectToGrpcServiceParams) []grpc.StreamClientInterceptor {
	streamClientInterceptors := baseStreamClientInterceptors

	if params.StreamClientInterceptors != nil {
		streamClientInterceptors = append(streamClientInterceptors, params.StreamClientInterceptors...)
	}

	// traceInterceptor := grpctrace.StreamClientInterceptor(grpctrace.WithServiceName(params.Name))
	// streamClientInterceptors = append(streamClientInterceptors, traceInterceptor)

	return streamClientInterceptors
}

// buildUnaryClientInterceptors returns the default list of interceptors for a unary client connection.
// Additional interceptors can be added to this list by passing in params.UnaryClientInterceptors.
func buildUnaryClientInterceptors(params ConnectToGrpcServiceParams) []grpc.UnaryClientInterceptor {
	unaryClientInterceptors := baseUnaryClientInterceptors

	if params.UnaryClientInterceptors != nil {
		unaryClientInterceptors = append(unaryClientInterceptors, params.UnaryClientInterceptors...)
	}

	// traceInterceptor := grpctrace.UnaryClientInterceptor(grpctrace.WithServiceName(params.Name))
	// unaryClientInterceptors = append(unaryClientInterceptors, traceInterceptor)

	return unaryClientInterceptors
}

// NewGrpcClientConn creates a standardized dial context, the connection is established in the background.
func NewGrpcClientConn(params ConnectToGrpcServiceParams) (*grpc.ClientConn, error) {
	if params.OnConnect != nil {
		return nil, errors.New("OnConnect parameter is supported")
	}

	dialOpts := newDialOptions(params)

	var clientConn *grpc.ClientConn
	var err error

	log.Info().Msgf("creating new grpc client connection to %s at %s", params.Name, params.Address)
	clientConn, err = grpc.DialContext(context.Background(), params.Address, dialOpts...)

	return clientConn, err
}
