package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/bluebell/controller"
	"github.com/sjxiang/bluebell/pkg/middlewares"
)

// 注册业务路由
func registerApiRoutes(router *gin.Engine) {

	// 用户注册
	router.POST("/signup", controller.SignUpHandler)

	// 用户登录
	router.POST("/login", controller.LoginHandler)

	// 身份验证
	router.GET("/validate", middlewares.JWTAuth, controller.ValidateHandler)
}