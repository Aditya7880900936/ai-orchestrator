package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter *rate.Limiter
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	c, exists := clients[ip]
	if !exists {
		// 10 requests/minute, burst 5
		limiter := rate.NewLimiter(rate.Every(6e9), 5)
		clients[ip] = &client{
			limiter: limiter,
		}
		return limiter
	}

	return c.limiter
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		limiter := getLimiter(c.ClientIP())

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
