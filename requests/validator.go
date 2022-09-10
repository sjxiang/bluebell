package requests



import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 校验参数函数类型
type ValidateFunc func(interface{}, *gin.Context) map[string][]string


func Validate(ctx *gin.Context, obj interface{}, handler ValidateFunc) bool {
	
	// 1. 解析请求，支持 JSON 数据
	if err := ctx.ShouldBindJSON(obj); err != nil {

		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{  // 422 语义错误 
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error": err.Error,
		})
		
		return false
	}

	// 2. 表单验证
	errs := handler(obj, ctx)  // 勾子
	
	// 3. 判断验证是否 pass
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{  // 422
			"message": "请求验证不通过，具体请查看 errors",
			"errors": errs,
		})

		return false
	} 

	return true
}


func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
