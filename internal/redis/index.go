package redis

import (
	"github.com/douyu/jupiter/pkg/client/redis"
	"github.com/douyu/jupiter/pkg/xlog"
	"time"
)

var RedisStub *redis.Redis

func InitRedis() error {

	RedisStub = redis.StdRedisStubConfig("kjcloud").Build()
	setRes := RedisStub.Set("jupiter-redis", "redisStub", time.Second*5)
	xlog.Info("redisStub set string", xlog.Any("res", setRes))

	getRes := RedisStub.Get("jupiter-redis")
	xlog.Info("redisStub get string", xlog.Any("res", getRes))
	return nil
}
