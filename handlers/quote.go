package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-redis-shell/models"
	"log"

	redis1 "gin-redis-shell/redis"
	"io"

	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// GetDailyQuote 获取每日一言的Http处理函数
func GetDailyQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//解析查询参数
	tag := r.URL.Query().Get("tag")
	name := r.URL.Query().Get("name")
	//从redis获取数据
	quote, err := redis1.GetQuote(tag, name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Quote not found"})
		return
	}
	//返回json响应
	json.NewEncoder(w).Encode(quote)
}

// 从网站 Api中获取每日一言
func GetDailyQuoteFromApi() (*models.QuoteResponse, error) {
	apiURL := "https://api.xygeng.cn/one" //c=a表示所有类型
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("http请求失败:%v", err)

	}
	defer resp.Body.Close()
	//检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API返回非200状态码:%d", resp.StatusCode)

	}
	//读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败，%v", err)
	}
	var quote models.QuoteResponse
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, fmt.Errorf("解析json失败,%v", err)
	}
	return &quote, nil
}

// 保存到Redis
func SaveToRedis(rdb *redis.Client, key string, data interface{}, expiration time.Duration) error {
	ctx := context.Background()
	if rdb == nil {
		return fmt.Errorf("redis client is nil")
	}
	if data == nil {
		return fmt.Errorf("data is nil")
	}
	//将数据序列化为json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json序列化失败:%v", err)
	}

	quote, err := GetDailyQuoteFromApi()
	if err != nil {
		log.Printf("获取每日语录失败：%v", err)
	} else {
		if quote == nil {
			key = "quote:daily"
		} else if quote.Data.Name != "" && quote.Data.Tag != "" {
			key = fmt.Sprintf("quote:%s:%s", quote.Data.Tag, quote.Data.Name)
		} else if quote.Data.Tag != "" {
			key = fmt.Sprintf("quote:tag:%s", quote.Data.Tag)
		} else if quote.Data.Name != "" {
			key = fmt.Sprintf("quote:name:%s", quote.Data.Name)
		} else {
			key = "quote:daily"
		}
	}
	//存储到Redis
	err = rdb.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("Redis存储失败:%v", err)
	}
	fmt.Println(err)
	log.Printf("数据已存入Redis,键名:%s", key)
	return nil
}
