package router

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

func (r *routerImpl) Shutdown(ctx context.Context) error {

	if r.corImpl == nil {
		return errors.New("the concierge service was nil when attempting to shutdown router")
	}

	r.grpcSvrShutdown()
	log.Info().Msg("GRPC servers stopped")
	r.app.Shutdown()
	log.Info().Msg("Router stopped")
	r.corImpl.Shutdown(ctx)
	log.Info().Msg("Concierge stopped")
	// r.contImpl.Shutdown(ctx)
	// log.Info().Msg("Content service stopped")
	// r.idImpl.Shutdown(ctx)
	// log.Info().Msg("Identity service stopped")
	r.finImpl.Shutdown(ctx)
	log.Info().Msg("Finance service stopped")
	// oauth.Close(ctx)

	// TODO: deprecate. this service is an internal on and can only be called via
	// if r.app == nil {
	// 	return errors.New("app is nil when attempting to shutdown router")
	// }
	// if err := r.app.Shutdown(); err != nil {
	// 	return err
	// }
	return nil
}
