package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ParamPostVote struct {
	// UserID 从请求中获取当前的用户
	PostID    int64  `json:"post_id,string"   valid:"post_id"`    // 帖子 id
	Direction int64  `json:"direction,string" valid:"direction"`  // 赞成票（1），反对票（-1）
}


func PostVote(data interface{}, ctx *gin.Context) map[string][]string {

	// 1. 自定义规则
	rules := govalidator.MapData{
		"post_id":   []string{"required", "digits:18"},
		"direction": []string{"required", "digits:2"},
	}

	// 2. 自定义验证出错时的提示
	messages := govalidator.MapData{
		"post_id": []string{
			"required:帖子编号为必填项",
			"digits:id 长度为 18",
		},
		"direction": []string{
			"required:支持或者反对为必填项",
			"digits:id 长度为 2",
		},
	}

	// 3. 开始验证
	errs := validate(data, rules, messages)

	return errs
}

