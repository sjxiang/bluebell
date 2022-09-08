package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/bluebell/logic"
	"github.com/sjxiang/bluebell/requests"
)


// SignUpHandler 处理注册请求
func SignUpHandler(ctx *gin.Context) {

	// 1. 获取请求参数 & 参数校验
	p := new(requests.ParamSignup)

	if ok := requests.Validate(ctx, p, requests.Signup); !ok {
		return
	}

	// 2. 业务逻辑处理
	if err := logic.Signup(p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "用户注册失败",
			"Error": err.Error(),
		})

		return
	}

	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "用户注册成功",
	})
}



// LoginHandler 处理登录请求
func LoginHandler(ctx *gin.Context) {

	// 1. 获取请求参数 & 参数校验
	p := new(requests.ParamLogin)

	if ok := requests.Validate(ctx, p, requests.Login); !ok {
		return
	}


	// 2. 业务逻辑处理
	if err := logic.Login(p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "用户登录失败，用户名或密码错误",
			"Error": err.Error(),
		})

		return
	}

	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "用户登录成功",
	})

}


// ValidateHandler 身份验证
func ValidateHandler(ctx *gin.Context) {
	
	// 1. 获取请求参数 & 参数校验
	
	// 2. 业务逻辑处理
	
	// 3. 返回响应
}