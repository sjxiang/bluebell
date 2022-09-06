package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/bluebell/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery())

	
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "pong",
		})
	})

	return r
}