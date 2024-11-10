package utils

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/valyala/fastjson"
)

var (
	// parser pool
	SessStore      *session.Store
	JsonParserPool *fastjson.ParserPool
)

func init() {
	if JsonParserPool == nil {
		JsonParserPool = new(fastjson.ParserPool)
	}
}

func HasCache(redis *redis.Storage, key string, cache interface{}) (map[string]interface{}, error) {
	if redis == nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "Redis server has gone away")
	}
	data, err := redis.Get(key)
	var cacheData map[string]interface{}
	if err == nil && len(data) > 0 {
		if err := json.Unmarshal([]byte(data), &cacheData); err != nil {
			return nil, err
		}
	}
	return cacheData, nil
}
func SaveCache(redis *redis.Storage, key string, cache interface{}, duration time.Duration) error {
	if redis == nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Redis server has gone away")
	}
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	if err := redis.Set(key, data, duration); err != nil {
		return err
	}
	return nil
}
