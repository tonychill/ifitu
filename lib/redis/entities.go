package redis

import (
	"context"
	"errors"

	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/tonychill/ifitu/lib/utils"

	r9 "github.com/redis/go-redis/v9"
)

const (
	field_tokenization = ",.<>{}[]\"':;!@#$%^&*()-+=~"
)

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

var ErrNil = errors.New("juvae_search: nil returned")

type Config struct {
	Server        string `envconfig:"REDIS"`
	RedisCa       string `envconfig:"REDIS_CA"`
	RedisUserCert string `envconfig:"REDIS_USER_CERT"`
	RedisUserKey  string `envconfig:"REDIS_USER_KEY"`
	Index         string
	MaxIdle       int    `envconfig:"REDIS_MAXIDLE"`
	MaxActive     int    `envconfig:"REDIS_MAXACTIVE"`
	IdleTimeout   int64  `envconfig:"REDIS_IDLETIMEOUT"`
	ConnURL       string `envconfig:"REDIS_CONN_URL"`
	ConnAddr      string `envconfig:"REDIS_CONN_ADDR"`
	UserName      string `envconfig:"REDIS_USERNAME"`
	Password      string `envconfig:"REDIS_PASSWORD"`
}

type Document struct {
	Id         string
	Score      float32
	Payload    []byte
	Properties map[string]interface{}
}

type Client interface {
	// This method will create an id if one was not provided in the request
	CreateJSON(ctx context.Context, req CreateJsonRequest) (string, error)
	GetJSON(ctx context.Context, query *global.Query, msg interface{}) ([]interface{}, error)
	JsonAppend(ctx context.Context, ap Append) error
	Do(ctx context.Context, args ...interface{}) *r9.Cmd
	Search(ctx context.Context, query *global.Query) (docs []Document, total int64, err error)
	SearchV2(ctx context.Context, query *global.Query, msg interface{}) (recs []interface{}, total int64, err error)
	HasIndex(ctx context.Context, index string) (bool, error)
}

type CreateJsonRequest struct {
	Id     string // TODO: deprecate: callers should not have to provide a key
	Prefix utils.Prefix
	Index  string
	Object any
	Args   []any
}
type clientImpl struct {
	name string
	conn *r9.Client
}

// Common filter
type Filter struct {
	Field   string
	Options interface{}
}

// FilterExpression the results to a given radius from lon and lat. Radius is given as a number and units
type GeoFilterOptions struct {
	Lon    float64
	Lat    float64
	Radius float64
	Unit   Unit
}

// units of Radius
type Unit string

const (
	KILOMETERS Unit = "km"
	METERS     Unit = "m"
	FEET       Unit = "ft"
	MILES      Unit = "mi"
)

// IndexingOptions represent the options for indexing a single document
type IndexingOptions struct {

	// If set, we use a stemmer for the supplied language during indexing. If set to "", we Default to English.
	Language string

	// If set to true, we will not save the actual document in the database and only index it.
	// As of RediSearch 2.0 and above NOSAVE is no longer supported, and will have no effect
	NoSave bool

	//  If set, we will do an UPSERT style insertion - and delete an older version of the document if it exists.
	Replace bool

	// (only applicable with Replace): If set, you do not have to specify all fields for reindexing.
	Partial bool

	// Applicable only in conjunction with Replace and optionally Partial
	// Update the document only if a boolean expression applies to the document before the update
	ReplaceCondition string
}

// DefaultIndexingOptions are the default options for document indexing
var DefaultIndexingOptions = IndexingOptions{
	Language:         "",
	NoSave:           false,
	Replace:          false,
	Partial:          false,
	ReplaceCondition: "",
}

// Argument is the interface implemented by an object which wants to control how
// the object is converted to Redis bulk strings.
type Argument interface {
	// RedisArg returns a value to be encoded as a bulk string per the
	// conversions listed in the section 'Executing Commands'.
	// Implementations should typically return a []byte or string.
	RedisArg() interface{}
}

// Limit results to those having numeric values ranging between min and max. min and max follow ZRANGE syntax, and can be -inf, +inf
type NumericFilterOptions struct {
	Min          float64
	ExclusiveMin bool
	Max          float64
	ExclusiveMax bool
}
