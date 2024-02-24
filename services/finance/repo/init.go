package repo

import (
	"context"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/lib/redis"
)

func InitializeRepository(ctx context.Context) (Repository, error) {
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
		conf := redis.Config{}
		if err := envconfig.Process("redis", &conf); err != nil {
			log.Fatal().Msgf("error when processing config from env: %s",
				err.Error())
		}
		r.redisConfig = conf
		ctx := context.Background()
		if err := r.initializeRedis(ctx); err != nil {
			log.Error().Err(err).Msg("failed to initialize finance service redis client")
			return nil, err
		}
		return r, nil
	}
}

func (r *repoImpl) initializeRedis(ctx context.Context) (err error) {
	r.redisClient, r.redisConfig, err = redis.NewClient("finance-service", os.Getenv("FINANCE_PAYMENTS"))
	if err != nil {
		return err
	}

	r.redisConfig.Index = os.Getenv("FINANCE_PAYMENTS")
	if err := r.createRatesIndex(ctx); err != nil {
		log.Error().Err(err).Msg("failed to verify or create index during finance initialization")
		return err
	}
	return nil
}
