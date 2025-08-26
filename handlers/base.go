package handlers

import (
	"gin-redis-shell/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseHandler struct {
}

func (*BaseHandler) Response(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusOK, dto.Resp{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
	} else {
		c.JSON(http.StatusOK, dto.Resp{
			Code: 200,
			Msg:  "成功",
			Data: data,
		})
	}
}
