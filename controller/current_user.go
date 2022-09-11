package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)


var ErrorUserNotLogin = errors.New("用户未登录")



func GetCurentUser(ctx *gin.Context) (int64, error) {
	
	uid, ok := ctx.Get("userID")

	if !ok {
		return 0, ErrorUserNotLogin
	}

	userID, ok := uid.(int64)
	if !ok {
		return 0, ErrorUserNotLogin
	}

	return userID, nil
}