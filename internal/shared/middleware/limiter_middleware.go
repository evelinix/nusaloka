package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var ctx = context.Background()

// RateLimitMiddleware membuat rate limiter per IP per endpoint
func RateLimitMiddleware(rdb *redis.Client, limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		endpoint := strings.ReplaceAll(c.FullPath(), "/", "_") // e.g. /auth/login => _auth_login
		key := "rl:" + ip + ":" + endpoint

		// increment counter di Redis
		val, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			log.Error().Err(err).Msg("Rate limiter error : ")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "rate limiter error"})
			return
		}

		// set TTL kalau belum ada
		if val == 1 {
			rdb.Expire(ctx, key, duration)
		}

		// kalau udah melewati limit, tolak request
		if int(val) > limit {
			log.Warn().Msgf("Rate limit exceeded for IP: %s, endpoint: %s", ip, endpoint)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Terlalu banyak permintaan, coba lagi nanti 🙅‍♀️",
			})
			return
		}

		// Optional logging for rate limits
		log.Info().Msgf("IP: %s, Endpoint: %s, Requests: %d", ip, endpoint, val)

		c.Next()
	}
}
