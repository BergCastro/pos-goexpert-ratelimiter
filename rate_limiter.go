package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	redisClient *redis.Client
	ipLimiter   map[string]*rate.Limiter
	tokenLimiter map[string]*rate.Limiter
	rateLimitIP  rate.Limit
	rateLimitToken rate.Limit
	blockTime    time.Duration
}

func NewRateLimiter(redisClient *redis.Client, rateLimitIP, rateLimitToken rate.Limit, blockTime time.Duration) *RateLimiter {
	return &RateLimiter{
		redisClient: redisClient,
		ipLimiter:   make(map[string]*rate.Limiter),
		tokenLimiter: make(map[string]*rate.Limiter),
		rateLimitIP:  rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockTime:    blockTime,
	}
}

func (rl *RateLimiter) getLimiter(key string, isToken bool) *rate.Limiter {
	if isToken {
		if limiter, exists := rl.tokenLimiter[key]; exists {
			return limiter
		}
		limiter := rate.NewLimiter(rl.rateLimitToken, int(rl.rateLimitToken))
		rl.tokenLimiter[key] = limiter
		return limiter
	} else {
		if limiter, exists := rl.ipLimiter[key]; exists {
			return limiter
		}
		limiter := rate.NewLimiter(rl.rateLimitIP, int(rl.rateLimitIP))
		rl.ipLimiter[key] = limiter
		return limiter
	}
}

func (rl *RateLimiter) Allow(key string, isToken bool) bool {
	limiter := rl.getLimiter(key, isToken)
	return limiter.Allow()
}

func (rl *RateLimiter) Block(key string, isToken bool) {
	ctx := context.Background()
	rl.redisClient.Set(ctx, key, "blocked", rl.blockTime)
}

func (rl *RateLimiter) IsBlocked(key string) bool {
	ctx := context.Background()
	val, err := rl.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		fmt.Println("Error checking block status:", err)
		return false
	}
	return val == "blocked"
}