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

	
	v1 := router.Group("/api/v1")

	v1.Use( middlewares.JWTAuth)  // 应用 JWT 认证中间件

	{

	}	

}