package clickup

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/tonychill/ifitu/lib/redis"
)

type Client interface {
	UpdatePaymentPaymentStatus(ctx context.Context, journeyId, experienceId string, status PaymentStatus) error
}

type clientImpl struct {
	apiKey      string
	baseUrl     string
	app         *fiber.App
	redisConfig redis.Config
	repo        *repo
}

var _ = Client(&clientImpl{})

type PaymentStatus struct {
}
type journey struct {
	// init_list_id:901700234454
}
type experience struct {
}
type note struct {
}
