package service

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"
	"github.com/tonychill/ifitu/services/finance/repo"
)

// InitOperationsService creates a new instance of the finance service
// TODO: add a logger and a connection to the database to the func params
func InitializeFinanceService(ctx context.Context) (*ServiceImpl, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	repo, err := repo.InitializeRepository(ctx)
	if err != nil {
		log.Error().Err(err).Msg("got an error when initializing the finance service repository")
		return nil, fmt.Errorf("failed to initiate finance repository: %w", err)
	}
	if repo == nil {
		return nil, fmt.Errorf("no repo was provided when attempting to initialize the [template] service")
	}

	return &ServiceImpl{
		repo: repo,
	}, nil
}
