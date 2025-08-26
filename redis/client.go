package redis1

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-redis-shell/models"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb  *redis.Client
	ctx  = context.Background()
	once sync.Once
)

// 初始化redis连接
func InitRedis(addr string, password string, db int) {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		//测试连接
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Connected to Redis successfully")
		}
	})

}

func GetRedisClient() *redis.Client {
	return rdb
}

// GetQuote 从Redis获取每日一言
func GetQuote(tag, name string) (*models.QuoteResponse, error) {
	var quote models.QuoteResponse
	var key string
	//根据查询条件构建Redis key
	if name != "" && tag != "" {
		key = fmt.Sprintf("quote:%s:%s", tag, name)
	} else if tag != "" {
		key = fmt.Sprintf("quote:tag:%s", tag)
	} else if name != "" {
		key = fmt.Sprintf("quote:name:%s", name)
	} else {
		key = "quote:daily"
	}
	//从redis获取数据
	data, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	//解析json数据
	err = json.Unmarshal([]byte(data), &quote)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}

// CacheQuote缓存每日一言到redis
func CacheQuote(quote *models.QuoteResponse, expiration time.Duration) error {
	//序列化为Json
	data, err := json.Marshal(quote)
	if err != nil {
		return err
	}
	//缓存主要key
	err = rdb.Set(ctx, "quote:daily", data, expiration).Err()
	if err != nil {
		return err
	}
	//缓存按tag索引
	for _, tag := range quote.Data.Tag {
		key := fmt.Sprintf("quote:tag:%v", tag)
		err = rdb.Set(ctx, key, data, expiration).Err()
		if err != nil {
			return err
		}
	}
	//缓存按作者名索引
	if quote.Data.Name != "" {
		key := fmt.Sprintf("quote:name:%s", quote.Data.Name)
		err = rdb.Set(ctx, key, data, expiration).Err()
		if err != nil {
			return err
		}
	}
	return nil

}
