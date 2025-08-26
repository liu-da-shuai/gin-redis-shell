package main

import (
	"fmt"
	"gin-redis-shell/config"
	"gin-redis-shell/router"
	"gin-redis-shell/spider"
)

//func main() {
//	//初始化Redis客户端
//
//	redisAddr := "localhost:6379"
//
//	redisPassword := ""
//	redisDB := 0
//	redis1.InitRedis(redisAddr, redisPassword, redisDB)
//	rdb := redis1.GetRedisClient()
//	fmt.Println(rdb)
//	//获取每日一言
//	quote, err := handlers.GetDailyQuoteFromApi()
//	if err != nil {
//		fmt.Println("获取每日一言失败")
//	}
//	fmt.Println(err)
//	//打印获取的每日一言 。。。
//	fmt.Printf("Got:%+v", quote)
//	fmt.Println("---------------------------------------")
//	//保存到redis，设置24小时过期
//
//	var key string
//	err = handlers.SaveToRedis(rdb, key, quote, 24*time.Hour)
//	if err != nil {
//		fmt.Printf("保存到Redis失败:%v", err)
//	} else {
//		fmt.Println("每日一言已成功获取并储存到Redis")
//	}
//	r := mux.NewRouter()
//	r.HandleFunc("/quote", handlers.GetDailyQuote).Methods("GET")
//	err = http.ListenAndServe(":8080", r)
//	if err != nil {
//		fmt.Printf("HTTP server start failed,err :%v", err)
//		return
//	}
//
//}

func main() {
	config.MockConfig() // to 后续用你yaml的来
	spider.NewSpider()
	if err := router.Start(); err != nil {
		fmt.Printf("HTTP server start failed,err :%v", err)
	}
	fmt.Println("HTTP server close")
}
