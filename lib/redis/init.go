package redis

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/kelseyhightower/envconfig"
	r9 "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// TODO: implement variadic func that will be used to create one or more indexes. Some
// services such as the concierge service will need to create multiple indexes.

type ClientOption func(*clientImpl) error

func NewClient(clientName, index string /*opts...add logger, indexingFunc func(ctx)error*/) (
	client Client, config Config, err error) {

	if index == "" {
		log.Fatal().Msg("index name cannot be empty")
	}
	config = Config{}
	if err := envconfig.Process("redis", &config); err != nil {
		log.Fatal().Msgf("error when processing Config from env: %s",
			err.Error())
	}
	config.Index = index
	if config.MaxIdle == 0 {
		config.MaxIdle = 1
	}
	if config.MaxActive == 0 {
		config.MaxActive = 10
	}
	if config.IdleTimeout == 0 {
		config.IdleTimeout = int64(240)
	}
	// log.Info().Msgf("has ca: %v, has cert: %v, has key: %v",
	// 	config.RedisCa != "", config.RedisUserCert != "", config.RedisUserKey != "")

	opts, err := r9.ParseURL(config.ConnURL)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse redis url")
		return nil, Config{}, err
	}
	opts.ClientName = clientName
	opts.MaxIdleConns = config.MaxIdle
	// opts := &r9.Options{
	// 	Network:      "tcp",
	// 	ClientName:   clientName,
	// 	Addr:         config.Server,
	// 	Password:     config.Password,
	// 	DB:           0,
	// 	MaxIdleConns: config.MaxIdle,
	// }

	if config.RedisUserCert != "" {
		// log.Info().Msgf("loading X509 cert and key pair")
		cert, err := tls.X509KeyPair([]byte(config.RedisUserCert), []byte(config.RedisUserKey))
		if err != nil {
			log.Error().Err(err).Msg("could not load redis keypair")
			return nil, Config{}, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(config.RedisCa))
		opts.TLSConfig = &tls.Config{
			RootCAs:            caCertPool,
			Certificates:       []tls.Certificate{cert},
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		}
		// log.Info().Msgf("setting pool for TLS enabled Redis server: %s, max idle: %d, max active: %d, idle timeout: %d",
		// 	config.Server, config.MaxIdle, config.MaxActive, config.IdleTimeout)

	} else {
		// log.Info().Msgf("setting pool for non TLS enabled Redis server: %s, max idle: %d, max active: %d, idle timeout: %d",
		// 	config.Server, config.MaxIdle, config.MaxActive, config.IdleTimeout)

	}

	return &clientImpl{
		name: index,
		conn: r9.NewClient(opts),
	}, config, nil
}
