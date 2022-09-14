package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"


	"github.com/sjxiang/bluebell/pkg/serializer"
	"github.com/sjxiang/bluebell/requests"
	"github.com/sjxiang/bluebell/logic"
)

// 投票
func PostVoteHandler(ctx *gin.Context) {
	
	// 参数校验
	p := new(requests.ParamPostVote)
	if ok := requests.Validate(ctx, p, requests.PostVote); !ok {
		return 
	}
	
	userID, err := GetCurentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, serializer.ParamErr("需要登陆", err))
		return
	}



	// 投票逻辑
	err = logic.PostVote(userID, p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.DBErr("", err))
	}

	ctx.JSON(http.StatusOK, serializer.Response{
		Data: p,
	})
}
