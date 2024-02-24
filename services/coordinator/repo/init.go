package repo

import (
	"context"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/lib/redis"
)

func InitiateRepository(_ context.Context) (Repository, error) {
	switch os.Getenv("ENV") {
	case "test":
		log.Debug().Msg("*************** using mock repo ***************")
		return newMockRepository()
	default:
		r := &repoImpl{}
		if err := envconfig.Process("", r); err != nil {
			log.Fatal().Msgf("error when processing config from env: %s",
				err.Error())
		}

		ctx := context.Background()
		if err := r.initializeRedis(ctx); err != nil {
			log.Error().Err(err).Msg("failed to initialize coordinator redis client")
			return nil, err
		}
		return r, nil
	}
}

func (r *repoImpl) initializeRedis(ctx context.Context) (err error) {
	r.redisClient, r.redisConfig, err = redis.NewClient("coordinator-service", "coordinator_service")
	if err != nil {
		return err
	}

	return nil
}
