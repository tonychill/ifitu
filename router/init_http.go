package router

import (
	"context"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

const (
	MB     = 1 << 20
	GB     = 1 << 30
	TB     = 1 << 40
	MaxAge = 86900 * 30
	IsProd = false
)

var (
	secretKey = os.Getenv("SESSION_SECRET_KEY")
)

func (r *routerImpl) initHttpServer(ctx context.Context, g *errgroup.Group) error {

	r.app = fiber.New(fiber.Config{
		BodyLimit:    100 * MB,
		ServerHeader: "tonychill/ifitu",
		AppName:      "coordinator_v1.0.1",
	})
	r.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, x-request-id, Authorization",
	}))

	if err := r.initOauth(ctx); err != nil {
		log.Error().Err(err).Msg("error initializing oauth")
		return err
	}

	/************ ROUTES **************/
	r.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the juvae!")
	})
	r.app.Static("/temp_images", "./temp_images")
	r.app.Post("/create-task", r.handleCreateTask)
	r.app.Post("/finance/stripe", r.handleStripeWebhook)
	r.app.Post("/add-payment", r.handleAddPayment)
	r.app.Post("/checkout", r.handleCheckout)
	r.app.Post("/webhook/clickup", r.handleCheckout)

	g.Go(func() error {
		if err := r.app.Listen(os.Getenv("SVC_HTTP_PORT")); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Could not start http server")
		}
		return nil
	})

	return nil
}

func (r *routerImpl) initOauth(ctx context.Context) error {

	// oauth.NewSessionStore(oauth.Config{
	// 	StoreProvider: oauth.FIBER,
	// 	// TODO: the following is pointless without integration with the identity service
	// 	// However, we still need to send the session info the the
	// 	// identity service so that it is in charge and can invalidate
	// 	// a guest if neccessary.
	// 	// Storage:        r.store,
	// })

	// googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	// googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// oauth.UseProviders(
	// 	google.New(googleClientId, googleClientSecret,
	// 		"http://localhost:4002/auth/callback/google",
	// 	),
	// )
	return nil
}
