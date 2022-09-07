package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/bluebell/middlewares"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(middlewares.Logger(), middlewares.Recovery())

	
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "pong",
		})
	})

	
	return r
}