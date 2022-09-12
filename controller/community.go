package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	
	"github.com/sjxiang/bluebell/pkg/serializer"
	"github.com/sjxiang/bluebell/logic"
	
)

// --- 跟社区论坛相关

func CommunityHandler(ctx *gin.Context) {
	
	
	// 查询到所有的社区（community_id、community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, serializer.Err(40001, "", err))  // 不轻易把服务端报错信息暴露给外面
		return
	}

	ctx.JSON(http.StatusOK, serializer.Response{
		Msg: "",
		Data: data,
	})

}