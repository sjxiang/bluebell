package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/bluebell/controller"
)

// 注册业务路由
func registerApiRoutes(router *gin.Engine) {

	// 用户注册
	router.POST("/signup", controller.SignUpHandler)

	// 用户登录

	// 用户校验
	
}