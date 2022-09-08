package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// "github.com/sjxiang/bluebell/middlewares"
)

// 路由初始化
func Setup(mode string) *gin.Engine {
	
	if mode == gin.ReleaseMode {  // "release"
		gin.SetMode(gin.ReleaseMode)
	} 
	
	router := gin.New()
	
	// 注册中间件 
	registerMiddleWare(router)

	// 注册业务路由
	registerApiRoutes(router)
	
	// 配置 ping 路由
	setupHealthCheckHandler(router)

	// 配置 404 路由
	setupNoFoundHandler(router)

	return router
}


func registerMiddleWare(router *gin.Engine) {
	// router.Use(
	// 	middlewares.Logger(), 
	// 	middlewares.Recovery(),
	// )
}


func setupNoFoundHandler(router *gin.Engine) {
	
	// 处理错误路由
	router.NoRoute(func(ctx *gin.Context) {

		// 获取 header 里面的 'Accept' 信息
		acceptStr := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptStr, "text/html") {

			// 如果是 HTML
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {

			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}


func setupHealthCheckHandler(router *gin.Engine) {

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "pong",
		})
	})
}