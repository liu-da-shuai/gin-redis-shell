package cache

import (
	"fmt"
	"gin-redis-shell/config"
	"gin-redis-shell/constant"
	"gin-redis-shell/database"
	"gin-redis-shell/dto"
	"gin-redis-shell/models"
	redis1 "gin-redis-shell/redis"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redis *redis.Client
	data  *database.QuoteModel
}

func NewCache() *Cache {
	addr := fmt.Sprintf("%s:%s", config.Conf.Redis.Host, config.Conf.Redis.Port)
	redis1.InitRedis(addr, config.Conf.Redis.Password, 0)
	return &Cache{
		redis: redis1.GetRedisClient(),
		data:  database.NewQuoteModel(),
	}
}

// todo
func (c *Cache) GetQuote(req *dto.QuoteReq) (*models.QuoteResponse, error) {
	data, err := redis1.GetQuote(req.Tag, req.Name) //
	if err != nil {
		return nil, err
	}
	if data.Data.Content == "" {
		// db 操作 有值回写 ，没有返回
		c.data.GetQuote(req.Tag, req.Name)
		// 回写 redis
		redis1.CacheQuote(data, constant.ExpireTime)
	}
	return data, nil
}
