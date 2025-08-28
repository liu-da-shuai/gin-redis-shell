package main

import (
	"gin-redis-shell/handlers"
	redis1 "gin-redis-shell/redis"
	"gin-redis-shell/router"
	"gin-redis-shell/config"
)

func main() {
	//初始化Redis客户端
	
	config.MockConfig()
	rdb := redis1.GetRedisClient()
	handlers.SaveAndGetQuote(rdb)
	router.Start()
}
