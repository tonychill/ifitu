package clickup

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/kelseyhightower/envconfig"
	"github.com/tonychill/ifitu/lib/redis"
)

type repo struct {
	redisClient      redis.Client
	JourneysIndex    string `envconfig:"CLICKUP_JOURNEYS" required:"true"`
	ExperiencesIndex string `envconfig:"CLICKUP_EXPERIENCES" required:"true"`
}

func (r *repo) initializeRepo(ctx context.Context) (err error) {
	if err := envconfig.Process("", r); err != nil {
		log.Error().Msgf("error when processing config from env: %s",
			err.Error())
	}
	r.redisClient, _, err = redis.NewClient("clickup-client", "clickup_client")
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) createSearchIndexes(ctx context.Context) (err error) {
	_, err = r.redisClient.Do(ctx, "FT.CREATE", r.JourneysIndex, "ON", "JSON",
		"PREFIX", "1", r.JourneysIndex+":", "SCHEMA",
		"$.journeyId", "as", "journeyId", "TEXT",
		"$.name", "as", "name", "TEXT",
	).Result()
	if err != nil {
		return err
	}
	_, err = r.redisClient.Do(ctx, "FT.CREATE", r.ExperiencesIndex, "ON", "JSON",
		"PREFIX", "1", r.ExperiencesIndex+":", "SCHEMA",
		"$.experienceId", "as", "experienceId", "TEXT",
		"$.name", "as", "name", "TEXT",
	).Result()
	if err != nil {
		return err
	}
	return nil
}
