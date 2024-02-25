package repo

import (
	"context"

	"github.com/rs/zerolog/log"
	// "gorm.io/gorm/logger"
)

// Creates search index for rates data.
func (r *repoImpl) createRatesIndex(ctx context.Context) (err error) {
	ok, err := r.redisClient.HasIndex(ctx, r.RatesIndex)
	if err != nil {
		log.Error().Err(err).Msgf("error checking if index %s exists", r.RatesIndex)
		return err
	}
	if ok {
		return nil
	}
	_, err = r.redisClient.Do(ctx, "FT.CREATE", r.RatesIndex, "ON", "JSON",
		"PREFIX", "1", r.RatesIndex+":", "SCHEMA",
		"$.id", "as", "id", "TEXT",
		"$.name", "as", "name", "TEXT",
		"$.experiences", "as", "experiences", "TEXT",
		"$.partnerId", "as", "partnerId", "TEXT",
		"$.rateType", "as", "rateType", "TEXT",
		"$.description", "as", "description", "TEXT",
		"$.currency", "as", "currency", "TEXT",
		"$.frequency", "as", "frequency", "TEXT",
		"$.createdAt", "as", "createdAt", "NUMERIC",
		"$.startDate", "as", "startDate", "NUMERIC",
		"$.updatedAt", "as", "updatedAt", "NUMERIC",
		"$.endDate", "as", "endDate", "NUMERIC",
		"$.amount", "as", "amount", "NUMERIC",
	).Result()
	if err != nil {
		log.Error().Err(err).Msgf("error creating index for finance rates with index name: %s",
			r.RatesIndex)
		return err
	}
	return nil
}
