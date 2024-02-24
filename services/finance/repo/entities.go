package repo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	finPb "github.com/tonychill/ifitu/apis/pb/go/finance"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/tonychill/ifitu/lib/redis"
	"github.com/tonychill/ifitu/lib/utils"
)

const (
	NO_END_EPOCH = 4072008207
)

type Repository interface {
	// CreateGuest(ctx context.Context, guest *global.Guest) error
	// GetGuests(ctx context.Context, query *global.Query) ([]*global.Guest, error)
	AddPaymentMethod(ctx context.Context, req *finSvc.AddPaymentMethodRequest) error
	GetPaymentMethods(ctx context.Context, req *finSvc.GetPaymentMethodsRequest) ([]*finPb.PaymentMethod, error)
	RemovePaymentMethod(ctx context.Context, req *finSvc.RemovePaymentMethodRequest) error
	Shutdown(ctx context.Context) error
}

var (
	_ = Repository(&repoImpl{})
)

type repoImpl struct {
	redisConfig redis.Config
	redisClient redis.Client
	shutdown    bool
	Env         string `envconfig:"ENV"`
	RatesIndex  string `envconfig:"FINANCE_PAYMENTS" requried:"true"`

	// pgClient            *pgClient
	// PgConnStr           string `envconfig:"POSTGRES_DB" required:"true"`
}

type internalRate struct {
	Id          string            `json:"id,omitempty"`
	PartnerId   string            `json:"partnerId,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Experiences []string          `json:"experiences,omitempty"`
	RateType    string            `json:"rateType,omitempty"`
	StartDate   int64             `json:"startDate,omitempty"`
	EndDate     int64             `json:"endDate,omitempty"`
	Amount      int64             `json:"amount,omitempty"`
	Currency    string            `json:"currency,omitempty"`
	Frequency   string            `json:"frequency,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   int64             `json:"createdAt,omitempty"`
	UpdatedAt   int64             `json:"updatedAt,omitempty"`
}

func (r *internalRate) fromProto(rp *global.Rate, new bool) {
	r.Id = rp.Id
	r.PartnerId = rp.PartnerId
	r.Name = rp.Name
	r.Description = rp.Description
	r.Experiences = rp.Experiences
	r.StartDate = rp.StartDate
	r.EndDate = rp.EndDate
	r.Amount = rp.Amount
	r.Metadata = rp.Metadata
	r.CreatedAt = rp.CreatedAt
	r.UpdatedAt = rp.UpdatedAt
	r.RateType = rp.RateType.String()
	r.Currency = rp.Currency.String()
	r.Frequency = rp.Frequency.String()
	if new {
		id, err := utils.NewULID(utils.Finance)
		if err != nil {
			log.Error().Err(err).Msgf("error generating ULID for Rate data: %+v", r)
			panic(err)
		}
		r.Id = id
		r.CreatedAt = time.Now().UTC().UnixMilli()
	}
}

func (r *internalRate) toProto() *global.Rate {
	return &global.Rate{
		Id:          r.Id,
		PartnerId:   r.PartnerId,
		Name:        r.Name,
		Description: r.Description,
		Experiences: r.Experiences,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
		Amount:      r.Amount,
		Metadata:    r.Metadata,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		RateType:    global.RateType(global.RateType_value[r.RateType]),
		Currency:    global.Currency(global.Currency_value[r.Currency]),
		Frequency:   global.RateFrequency(global.RateFrequency_value[r.Frequency]),
	}

}
