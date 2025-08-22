package main

import (
	"fmt"
	"net/http"
	"time"

	"gin-redis-shell/handlers"
	redis1 "gin-redis-shell/redis"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	//初始化Redis客户端

	redisAddr := "localhost:6379"

	redisPassword := ""
	redisDB := 0
	redis1.InitRedis(redisAddr, redisPassword, redisDB)
	fmt.Println(rdb)
	//获取每日一言
	quote, err := handlers.GetDailyQuoteFromApi()
	if err != nil {
		fmt.Println("获取每日一言失败")
	}
	fmt.Println(err)
	//打印获取的每日一言 。。。
	fmt.Printf("Got:%q", quote)
	fmt.Println("---------------------------------------")
	//保存到redis，设置24小时过期

	var key string
	err = handlers.SaveToRedis(rdb, key, quote, 24*time.Hour)
	if err != nil {
		fmt.Printf("保存到Redis失败:%v", err)
	} else {
		fmt.Println("每日一言已成功获取并储存到Redis")
	}
	r := mux.NewRouter()
	r.HandleFunc("/quote", handlers.GetDailyQuote).Methods("GET")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err :%v", err)
		return
	}

}
