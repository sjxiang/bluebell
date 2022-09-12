package requests


import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)


// 创建帖子请求参数
type ParamCreatePost struct {
	Title   	  string `json:"title"         valid:"title"`
	Content       string `json:"content"       valid:"content"` 
	CommunityID   string `json:"community_id"  valid:"community_id"` // 社区话题分类，前端下拉框提供
}



func CreatePost(data interface{}, ctx *gin.Context) map[string][]string {

	// 1. 自定义规则
	rules := govalidator.MapData{
		"title":        []string{"required", "between:3,20"},
		"content":      []string{"required", "between:10,1024"},
		"community_id": []string{"required", "digits:6"},		
	}

	// 2. 自定义验证出错时的提示
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"between:标题长度需在 3~20 之间",
		},
		"content": []string{
			"required:内容为必填项",
			"between:内容长度需在 10~1024 之间",
		},
		"community_id": []string{
			"required:社区话题为必填项",
			"digits:社区话题长度必须为 6 位的数字",
		},
	}

	// 3. 开始验证
	errs := validate(data, rules, messages)
	
	return errs
}