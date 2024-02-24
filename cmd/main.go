package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/tonychill/ifitu/router"

	// cont "github.com/tonychill/ifitu/services/content/service"
	cor "github.com/tonychill/ifitu/services/coordinator/service"
	fin "github.com/tonychill/ifitu/services/finance/service"

	// iden "github.com/tonychill/ifitu/services/identity/service"

	// opsSvc "github.com/tonychill/ifitu/services/journeys/service"
	// coorSvc "github.com/tonychill/ifitu/services/journeys/service"
	"golang.org/x/sync/errgroup"
)

func main() {

	ctx := context.Background()
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Setup signal interuption for graceful shutdowns
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// This will be used to connect to the databases. those connections
	// will be passed to the individual services so that they are able
	// use those connections when instantiating their repository imiplimentations
	// _ = connectToDBs(ctx)

	// contSvc, err := cont.InitializeContentService(ctx)
	// if err != nil {
	// 	log.Fatal().Msgf("failed to initialize content service: %v", err)
	// }

	// log.Info().Msgf("content service initialized\n")

	// idSvc, err := iden.InitializeIdentityService(ctx)
	// if err != nil {
	// 	log.Fatal().Msgf("failed to initialize identity service: %v", err)
	// }

	// log.Info().Msgf("identity service initialized\n")

	finSvc, err := fin.InitializeFinanceService(ctx)
	if err != nil {
		log.Fatal().Msgf("failed to initialize finance service: %v", err)
	}

	log.Info().Msgf("finance service initialized\n")

	conSvc, err := cor.InitializeConciergeService(ctx)
	if err != nil {
		log.Fatal().Msgf("failed to initialize coordinator: %v", err)
	}

	log.Info().Msgf("coordinator initialized\n")

	// TODO: implement a way to validate that when a new service is added
	// that they are initialized correctly and that they are added to the
	// router's service implementations and completely registered throughout.
	rtr, err := router.InitializeRouter(ctx, g, router.ServiceImplementations{
		CoordinatorService: conSvc,
		// ContentService:   contSvc,
		// IdentityService:  idSvc,
		FinanceService: finSvc,
	})
	if err != nil {
		log.Fatal().Msgf("failed to initialize router: %v", err)
	}
	log.Info().Msgf("router initialized with service implementations\n")
	log.Info().Msg("🎉🎉🎉 Welcome to the good vibes 🎉🎉🎉")

	// Handle shutdown from a signal
	select {
	case <-interrupt:
		log.Info().Msgf("Received interrupt signal, shutting down...")
		break
	case <-ctx.Done():
		break
	}

	if err := rtr.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("error shutting down after service interrupt: %v", err)
	}

	_, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	// Wait on the goroutines to finish
	log.Info().Msg("waiting for goroutines to finish...")
	if err := g.Wait(); err != nil {
		log.Error().Err(err).Msg("error waiting for goroutines to finish")
		os.Exit(2)
	}
}
