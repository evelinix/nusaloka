package middleware

import (
	"log"
	"net/http"

	libredis "github.com/redis/go-redis/v9"
	limiter "github.com/ulule/limiter/v3"
	mhttp "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var RateLimitHandler func(http.Handler) http.Handler

func InitRateLimiter() error {
	// Define rate: 10 req per minute
	rate, err := limiter.NewRateFromFormatted("10-M")
	if err != nil {
		return err
	}

	// Create a redis client.
	option, err := libredis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatal(err)
		return err
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_chi_example",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create a new middleware with the limiter instance.
	rateMiddleware := mhttp.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))

	RateLimitHandler = rateMiddleware.Handler
	return nil
}
