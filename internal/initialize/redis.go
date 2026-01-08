package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/pkg/setting"
	"github.com/anhnv1202/base-go/pkg/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	m := global.Config.Redis

	if m.IsSentinelMode() {
		global.Redis = createSentinelClient(m)
		global.Logger.Info("Redis sentinel mode", zap.String("master", m.MasterName))
	} else {
		global.Redis = createStandaloneClient(m)
		global.Logger.Info("Redis standalone mode", zap.String("addr", m.Addr()))
	}

	healthCheckRedis()

	OnShutdown(func() {
		if global.Redis != nil {
			global.Redis.Close()
			global.Logger.Info("Redis connection closed")
		}
	})
}

func createStandaloneClient(m setting.RedisConfig) redis.UniversalClient {
	return redis.NewClient(&redis.Options{
		Addr:         m.Addr(),
		Password:     m.Password,
		DB:           m.DB,
		PoolSize:     m.PoolSize,
		MinIdleConns: m.MinIdleConns,
		MaxRetries:   m.MaxRetries,
		DialTimeout:  m.DialTimeout,
		ReadTimeout:  m.ReadTimeout,
		WriteTimeout: m.WriteTimeout,
		PoolTimeout:  m.PoolTimeout,
	})
}

func createSentinelClient(m setting.RedisConfig) redis.UniversalClient {
	addrs := m.GetSentinelAddrs()
	if len(addrs) == 0 {
		utils.CheckErrorPanic(fmt.Errorf("no sentinel addresses configured"), "REDIS_SENTINEL_ADDRS required")
	}
	if m.MasterName == "" {
		utils.CheckErrorPanic(fmt.Errorf("master name not configured"), "REDIS_MASTER_NAME required")
	}

	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       m.MasterName,
		SentinelAddrs:    addrs,
		SentinelPassword: m.SentinelPassword,
		Password:         m.Password,
		DB:               m.DB,
		PoolSize:         m.PoolSize,
		MinIdleConns:     m.MinIdleConns,
		MaxRetries:       m.MaxRetries,
		DialTimeout:      m.DialTimeout,
		ReadTimeout:      m.ReadTimeout,
		WriteTimeout:     m.WriteTimeout,
		PoolTimeout:      m.PoolTimeout,
	})
}

func healthCheckRedis() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	utils.CheckErrorPanic(global.Redis.Ping(ctx).Err(), "failed to ping Redis")

	s := global.Redis.PoolStats()
	global.Logger.Info("Redis connected",
		zap.Uint32("conns", s.TotalConns),
		zap.Uint32("idle", s.IdleConns),
	)
}
