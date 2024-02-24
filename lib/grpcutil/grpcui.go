package grpcutil

import (
	"context"
	"net/http"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCUIServer returns a http server with a single admin route. This route is served by grpcui's web handler, which
// allows invoking gRPC sever methods. The gRPC server must enable server reflection.
func newGRPCUIServer(ctx context.Context, httpListen, grpcListen string) (*http.Server, func(context.Context) []error, error) {
	cc, err := grpc.DialContext(ctx, grpcListen, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	adminHandler, err := standalone.HandlerViaReflection(ctx, cc, grpcListen)
	if err != nil {
		return nil, nil, err
	}

	adminRouter := mux.NewRouter()

	// grpc-ui requires serving content from the "/" path, so rewrite /_admin or /_admin/ to /
	adminRouter.PathPrefix("/_admin/").Handler(http.StripPrefix("/_admin", adminHandler))
	adminRouter.Handle("/_admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/_admin/", http.StatusMovedPermanently)
	}))

	adminServer := http.Server{
		Addr:    httpListen,
		Handler: adminRouter,
	}

	tearDown := func(ctx context.Context) []error {
		var errors []error
		if err = cc.Close(); err != nil {
			errors = append(errors, err)
		}
		if err := adminServer.Shutdown(ctx); err != nil {
			errors = append(errors, err)
		}
		return errors
	}

	return &adminServer, tearDown, nil
}
