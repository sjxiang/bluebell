package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/bluebell/pkg/jwt"
	"github.com/sjxiang/bluebell/pkg/serializer"
)


// 基于 JWT 的认证中间件
func JWTAuth(ctx *gin.Context) {
	
	// 客户端携带 Token 有三种方式，具体取决于实际业务情况
	// 1. 放在请求 header
	// 2. 放在请求 body 
	// 3. 放在 url 中
	//
	// 方式一，具体做法：
	// token 放在 Header 的 Authorization 中，例如 "bearer xxx.xxx.xxx"
	
	authHeader := ctx.Request.Header.Get("Authorization")
	
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, serializer.Response{
			Code: 401,
			Msg: "请求 header 中 auth 为空",
		})
		ctx.Abort()
		return
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ctx.JSON(http.StatusUnauthorized, serializer.Response{
			Code: 401,  // 格式
			Msg: "请求 header 中，auth 格式有误",
		})
	
		ctx.Abort()
		return
	}	

	// parts[1] 是获取到的 tokenString，可以用 jwt 包方法解析
	claims, err := jwt.NewJWT().ParseToken(parts[1]); 
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, serializer.Response{
			Code: 401,
			Msg: "无效的 token",
			Error: err.Error(),
		})

		ctx.Abort()
		return
	}

	// 检查 expired time
	if float64(time.Now().Unix()) > float64(claims.ExpiresAt) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	

	// TODO 校验
	
	// 将当前请求的 userID 信息保存到请求的上下文 ctx 中
	ctx.Set("userID", claims.UserID)
	ctx.Next()  
	
	// 后续的 handler 可以用 ctx.Get("userID") 来获取当前请求的用户信息

}