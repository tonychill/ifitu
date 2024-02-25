package repo

import (
	"context"

	"github.com/tonychill/ifitu/lib/redis"
)

type Repository interface {
	// GetGuests(ctx context.Context, query *global.Query) ([]*identity.Guest, error)
	// UpdateGuest(ctx context.Context, guest *identity.Guest) ([]*identity.Guest, error)
	// CreateRelationship(ctx context.Context, rel Relationship) error

	Shutdown(ctx context.Context) error

	// Deprecated
	// __SaveFlow(ctx context.Context, req *pb.Flow) (string, error)
	// __GetFlowStatus(ctx context.Context, req *pb.CheckFlowStatusRequest) ([]*pb.FlowStatus, error)
	// __GetItems(context.Context, *global.Query) ([]*concierge.Item, error)
}
type Relationship struct {
	Id         string
	Name       string
	Type       string
	Properties map[string]string
}

var (
	_ = Repository(&repoImpl{})
)

type repoImpl struct {
	redisConfig     redis.Config
	redisClient     redis.Client
	shutdown        bool
	Env             string `envconfig:"ENV"`
	GraphDbAddress  string `envconfig:"GRAPH_DB_ADDRESS" required:"false"` // FIXME: required, false for now
	FlowIndexPrefix string `envconfig:"COORDINATOR_FLOWS" required:"true"`

	// pgClient            *pgClient
	// PgConnStr           string `envconfig:"POSTGRES_DB" required:"true"`
	// FIXME: deprecate in favor of using the consumer group for tracking
}

// 2-25T02:54:12.235