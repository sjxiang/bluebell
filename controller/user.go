package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/bluebell/logic"
	"github.com/sjxiang/bluebell/requests"
)

// SignUpHandler 处理注册请求
func SignUpHandler(ctx *gin.Context) {

	// 1. 获取参数 & 参数校验
	p := new(requests.ParamSignup)

	if ok := requests.Validate(ctx, p, requests.Signup); !ok {
		return
	}

	// 2. 业务处理
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

