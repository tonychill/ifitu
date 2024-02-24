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
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		BodyLimit:    100 * MB,
		ServerHeader: "tonychill/ifitu",
		AppName:      "coordinator_v1.0.1",
	})
	r.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, x-request-id, Authorization",
	}))

	// store := sessions.NewCookieStore([]byte(secretKey))
	// store.MaxAge(MaxAge)
	// store.Options.Path = "/"
	// store.Options.Secure = IsProd
	// store.Options.HttpOnly = true
	// b, _ := json.Marshal(store)
	// fmt.Printf("***** TESTING: json string of store: %s", b)
	// store.Config.CookieDomain
	// ctx := context.Background()
	// if err := r.conSvc.InitAuth(ctx, idImpl.InitAuthRequest{}); err != nil {
	// 	panic("couldn't initiat auth with store")
	// }

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

	// TODO: Add a way to get all the methods from the service interface
	// and create the routes automatically
	// t := reflect.TypeOf((*service.Service)(nil)).Elem()
	// var s []string
	// for i := 0; i < t.NumMethod(); i++ {
	// 	s = append(s, t.Method(i).Name)
	// }

	/************ ROUTES **************/
	r.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the juvae!")
	})

	r.app.Static("/temp_images", "./temp_images")
	/*

			guest inputs the name and type
			makes a call to getRules() -> rules : select [rules]
			if len(rules) == 0 : call create createRules()
			if rule req resource
				getResources() - [resournces]
				if len(resources) == 0
					createResournce() -> {id,...}
				addResourceToRule()
			create-experience(
			experience,
			rules
			:: rules will tie the experience to the resource(s)
			:: and must be created before the experience can be
			:: created as the rules are essentially what enables
			:: the exisitence of the experiences
			-> get-resources()->resources[] | create-resource()->id:string
			)

			only once the rule is selected/created will the experience be created.
			-- get-rules

			content uploaded with the resource data
			create-resource(
				resource,
				content
			)
					add ruled
						if rule requires a resource (e.g. boat)
							{id,} <-
							[ids] <- upload-content(ownerId, resourceId, content)

					if experience has content, upload content to cloudflare

					/*
		   1. create experience takes the initial data sucha s name, description, etc... and saves them to localstate (localStorage too??)
		   2. next step lets the partner create or select the rules associated with the experience; pricing, req resources, availability, max pax, etc..
		      a. if a rule does not exist then the client will go to the rule creation step (create-rule)
		      b. if rule requires a non-existing resource then the client will go to the create resource step (create-resource)
		   3. finally, when the partner clicks done, the client will then call 'create-experience' containing all the data related to the experience such as the rules and the required/optional resources.
	*/
	// r.app.Get("/login/:provider", r.handleLogin)
	// r.app.Get("/auth/callback/:provider", r.handleAuthCallback)
	// r.app.Get("/logout", r.handleLogout)
	// r.app.Get("get-session", r.handleGetSession)
	// r.app.Post("/create-guest", r.handleCreateGuest)
	r.app.Post("/create-task", r.handleCreateTask)

	// TODO: In work
	r.app.Post("/finance/stripe", r.handleStripeWebhook)

	// TODO: Deprecating
	// r.app.Post("/initiate-flow", r.handleInitiateFlow)
	// r.app.Post("/get-resources", r.handleGetResources)
	// r.app.Delete("/delete/:image", nil /*handleDeleteImage*/)

	g.Go(func() error {
		if err := r.app.Listen(os.Getenv("SVC_HTTP_PORT")); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Could not start http server")
		}
		return nil
	})

	return nil
}
