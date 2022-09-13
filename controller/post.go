package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/bluebell/logic"
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

	userID, err := GetCurentUser(ctx) 
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, serializer.DBErr("", err))

		return
	}

	// 2. 创建帖子
	if err := logic.CreatePost(p, userID); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, serializer.DBErr("", err))

		return
	}


	// 3. 返回响应	
	ctx.JSON(http.StatusOK, serializer.Response{
		Msg: "创建 post 成功",
	})
	
}


// 查看帖子详情
func GetPostDetailHandler(ctx *gin.Context) {

	// 1. 获取帖子的 id
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}


	// 2. 根据 id 取出帖子数据
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic GetPostByID failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, serializer.DBErr("", err))
		return
	}

	// 3. 返回响应
	ctx.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildPost(*data),
	})
}



// 分页查询帖子
func GetPostListHandler(ctx *gin.Context) {

	// 获取分页查询参数
	pageSizeStr := ctx.Query("pageSize")
	pageNumStr := ctx.Query("pageNum")

	
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, serializer.ParamErr("偏移量参数错误", err))
		return
	}

	pageNum, err := strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, serializer.ParamErr("限制参数错误", err))
		return
	}



	// 数据库查询
	data, err := logic.GetPostList(pageSize, pageNum)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.DBErr("", err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildPostList(data),
	})
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



