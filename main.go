package main

import (
	"gin-redis-shell/handlers"
	redis1 "gin-redis-shell/redis"
	"gin-redis-shell/router"
)

func main() {
	//初始化Redis客户端
	redisAddr := "localhost:6379"
	redisPassword := ""
	redisDB := 0
	redis1.InitRedis(redisAddr, redisPassword, redisDB)
	rdb := redis1.GetRedisClient()
	handlers.SaveAndGetQuote(rdb)
	router.Start()
}
