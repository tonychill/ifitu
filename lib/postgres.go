package lib

import (
	"context"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgClient struct {
	db *gorm.DB
}

type PgParams struct {
	ConnStr string
	Tables  []string
	Indexes []string
}

func NewPgClient(ctx context.Context, params *PgParams) (c *PgClient, err error) {
	c = &PgClient{
		// logger: params.logger,
	}

	// logger := params.logger.With().Context(ctx).Logger()
	params.ConnStr = "host=localhost user=tony password=password dbname=journeys port=5432 sslmode=disable"
	c.db, err = gorm.Open(postgres.Open(params.ConnStr), &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		// logger.Error().Err(err).Send()
		return nil, err
	}
	// TODO: loop through params.Tables and params.Indexes and create them if they don't exist
	if err := c.db.Exec(`
		CREATE TABLE IF NOT EXISTS 
			journeys (
				id						text				PRIMARY KEY,
			);
		`,
	).Error; err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			// logger.Error().Err(err).Msg("failure to create journeys table")
			return nil, err
		}
	}

	if err := c.db.Exec(`
		CREATE INDEX 
			journeys_idx
		ON journeys (guest_id);
		`,
	).Error; err != nil {
		// logger.Error().Err(err).Msg("failure to create journeys table")
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
	}

	return c, nil
}
