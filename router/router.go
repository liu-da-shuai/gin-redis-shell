package router

import (
	"fmt"
	"gin-redis-shell/config"
	"gin-redis-shell/handlers"
	"github.com/gin-gonic/gin"
)

func Start() error {
	engine := gin.Default()
	h := handlers.NewQuoteHandlers()
	engine.GET("/", h.GetDailyQuote)
	return engine.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))
}
