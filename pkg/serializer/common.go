package serializer

import "github.com/gin-gonic/gin"

// 基础序列化器

type Response struct {
	Code  int          `json:"code"`
	Msg   string       `json:"msg"`
	Error string       `json:"error,omitempty"`
	Data  interface{}  `json:"data,omitempty"`
}


// Err 通用错误
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg: msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}

	return res
}


// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "各种参数错误"
	}

	return Err(CodeParamErr, msg, err)
}