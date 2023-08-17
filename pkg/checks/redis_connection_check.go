package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/redis"
)

type RedisConnectionCheck struct {
	Config config.Config
	Label  string
}

// NewRedisConnectionCheck creates a new RedisConnectionCheck
func NewRedisConnectionCheck(config *config.Config, label string) *RedisConnectionCheck {
	return &RedisConnectionCheck{
		Config: *config,
		Label:  label,
	}
}

// Run runs the RedisConnectionCheck
func (r *RedisConnectionCheck) Run() error {
	redisClient := redis.NewRedisClient(&r.Config)
	if err := redisClient.Connect(); err != nil {
		return err
	}

	if err := redisClient.Ping(); err != nil {
		return err
	}

	return nil
}

// GetLabel returns the label of the RedisConnectionCheck
func (r *RedisConnectionCheck) GetLabel() string {
	return r.Label
}
