package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		clientIP := c.ClientIP()

		key := fmt.Sprintf("rate_limit:%s", clientIP)

		currentCount, err := redisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			err = redisClient.Set(ctx, key, 1, time.Minute).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
				c.Abort()
				return
			}
			redisClient.Expire(ctx, key, time.Minute)
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			c.Abort()
			return
		} else {
			count, _ := strconv.Atoi(currentCount)
			if count >= 10 {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "Rate limit exceeded. Try again later.",
				})
				c.Abort()
				return
			}

			redisClient.Incr(ctx, key)
		}

		c.Next()
	}
}
