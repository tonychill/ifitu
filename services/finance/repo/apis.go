package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	finPb "github.com/tonychill/ifitu/apis/pb/go/finance"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/tonychill/ifitu/lib/redis"
	"github.com/tonychill/ifitu/lib/utils"
)

func (r *repoImpl) CreateCustomer(ctx context.Context, guest *global.Customer) error {
	return nil
}

func (r *repoImpl) GetCustomers(ctx context.Context, query *global.Query) (guests []*global.Customer, err error) {
	recs, _, err := r.redisClient.SearchV2(ctx, query, &global.Customer{})
	if err != nil {
		log.Error().Err(err).Msg("error getting customer json from finance repository")
		return nil, err
	}
	for _, rec := range recs {
		g, ok := rec.(*global.Customer)
		if ok {
			guests = append(guests, g)
		}
	}
	if len(guests) == 0 {
		log.Error().Msgf("no guests found for this query: %v", query)
		return nil, fmt.Errorf("no guests were found with the given query")
	}

	return guests, nil
}

func (r *repoImpl) AddPaymentMethod(ctx context.Context, req *finSvc.AddPaymentMethodRequest) error {
	_, err := r.redisClient.CreateJSON(ctx, redis.CreateJsonRequest{
		Prefix: utils.PaymentMethod,
		Object: req.PaymentMethod,
	})
	if err != nil {
		log.Error().Err(err).Msgf("failed to add payemnt method for guest id %s", req.PaymentMethod.CustomerId)
		return err
	}
	return nil
}
func (r *repoImpl) GetPaymentMethods(ctx context.Context, req *finSvc.GetPaymentMethodsRequest) ([]*finPb.PaymentMethod, error) {
	docs, _, err := r.redisClient.Search(ctx, req.Query)
	if err != nil {
		log.Error().Err(err).Msg("Could not execute search")
		return nil, err
	}

	for _, doc := range docs {
		rate := &internalRate{}
		if err := doc.Decode(rate); err != nil {
			log.Error().Err(err).Msg("error converting redis document to rates to proto.")
			return nil, err
		}
	}

	return nil, nil
}
func (r *repoImpl) RemovePaymentMethod(ctx context.Context, req *finSvc.RemovePaymentMethodRequest) error {
	panic("implement me")
}

func (r *repoImpl) CreateRate(ctx context.Context, rate *global.Rate) (string, error) {
	return r.createRate(ctx, rate)
}
func (r *repoImpl) GetRates(ctx context.Context, query *global.Query) ([]*global.Rate, error) {
	return r.getRates(ctx, query)
}

func (r *repoImpl) getRates(ctx context.Context, query *global.Query) (rates []*global.Rate, err error) {
	docs, _, err := r.redisClient.Search(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Could not execute search")
		return nil, err
	}

	for _, doc := range docs {
		rate := &internalRate{}
		if err := doc.Decode(rate); err != nil {
			log.Error().Err(err).Msg("error converting redis document to rates to proto.")
			return nil, err
		}
		rates = append(rates, rate.toProto())
	}
	return
}

func (r *repoImpl) createRate(ctx context.Context, rate *global.Rate) (string, error) {

	if rate.EndDate == -1 {
		rate.EndDate = NO_END_EPOCH
	} else if time.Now().UTC().After(time.Unix(rate.EndDate, 0).UTC()) {
		log.Error().Msgf("the end date for %s rate is in the past: %d", rate.Name, rate.EndDate)
		return "", fmt.Errorf("the end date for the %s rate is in the past: %d", rate.Name, rate.EndDate)
	}

	ir := &internalRate{}
	ir.fromProto(rate, true)
	rate.Id = ir.Id

	if _, err := r.redisClient.CreateJSON(ctx, redis.CreateJsonRequest{
		Id:     ir.Id,
		Object: ir,
	}); err != nil {
		log.Error().Err(err).Msgf("error setting Journey data in redis for rate id %s", ir.Id)
		return "", err
	}
	return ir.Id, nil
}

func (r *repoImpl) Shutdown(ctx context.Context) error {
	r.shutdown = true
	// TODO: flush redis and other db connections.
	return nil
}
