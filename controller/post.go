package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/bluebell/pkg/serializer"
	"github.com/sjxiang/bluebell/requests"
)

// 增
func CreatePostHandler(ctx *gin.Context) {
	
	// 1. 获取请求参数 & 参数校验
	p := new(requests.ParamCreatePost)
	
	if ok := requests.Validate(ctx, p, requests.CreatePost); !ok {
		return
	} 

	ctx.JSON(http.StatusOK, serializer.Response{
		Data: &p,
	})
	// 2. 创建帖子
	
	// 3. 返回响应	

}


// 删
func DeletePostHandler(ctx *gin.Context) {
	// 1. 获取请求参数 & 参数校验
		
	// 2. 业务逻辑处理
		
	// 3. 返回响应	
}


// 改
func UpdatePostHandler(ctx *gin.Context) {
	// 1. 获取请求参数 & 参数校验
		
	// 2. 业务逻辑处理
		
	// 3. 返回响应	
}


// 查
func QueryPostHandler(ctx *gin.Context) {
	// 1. 获取请求参数 & 参数校验
		
	// 2. 业务逻辑处理
		
	// 3. 返回响应	
}



