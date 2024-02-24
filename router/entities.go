package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	// contImpl "github.com/tonychill/ifitu/services/content/service"
	corImpl "github.com/tonychill/ifitu/services/coordinator/service"
	finImpl "github.com/tonychill/ifitu/services/finance/service"
	// idImpl "github.com/tonychill/ifitu/services/identity/service"
)

type Router interface {
	Shutdown(ctx context.Context) error
}

type _Store struct {
	fs *session.Session
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}
type ServiceImplementations struct {
	CoordinatorService *corImpl.ServiceImpl
	// ContentService    *contImpl.ServiceImpl
	// IdentityService   *idImpl.ServiceImpl
	FinanceService *finImpl.ServiceImpl
}

// Provides a landing zone for incoming requests. The router will route requests to the
// concierge and then from there to the appropriate service(s).
type routerImpl struct {
	grpcSvrShutdown func()
	app             *fiber.App
	GrpcServerPort  string `envconfig:"SVC_GRPC_PORT" required:"true"`
	corImpl         *corImpl.ServiceImpl

	// The following services are here only so that they're registered by the router
	// and therefor can be called by the concierge. They should not be called directly.
	// This allows them to be called by the concierge. While it is possible for the
	// router to allow external callers to call the services independantly, it is not recommended.
	// The router's main purpose is to route requests to the concierge. The concierge will then
	// call the domain services as needed.
	// contImpl *contImpl.ServiceImpl
	// idImpl   *idImpl.ServiceImpl
	finImpl *finImpl.ServiceImpl
}
