package redis

import (
	"context"
	"fmt"
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	Host     string
	Password string
	Port     int
	Client   redis.Client
}

// NewRedisClient creates a new RedisClient
func NewRedisClient(config *config.Config) *RedisClient {
	return &RedisClient{
		Host:     config.Redis.Host,
		Password: config.Redis.Password,
		Port:     config.Redis.Port,
	}
}

// Connect connects to the Redis server
func (r *RedisClient) Connect() error {
	serverAddress := fmt.Sprintf("%s:%d", r.Host, r.Port)
	// create redis client options
	options := redis.Options{Addr: serverAddress}
	if r.Password != "" {
		options.Password = r.Password
	}

	// Create redis client.
	client := redis.NewClient(&options)
	r.Client = *client
	return nil
}

// Ping pings the Redis server
func (r *RedisClient) Ping() error {
	// Create a context with timeout for the connection check
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the Redis server
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
