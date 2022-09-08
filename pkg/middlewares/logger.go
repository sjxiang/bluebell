package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 记录访问日志
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		// 获取 resp content
		w := &responseBodyWriter{
			body: &bytes.Buffer{},
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = w

		// 获取 req 数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			
			// ctx.Request.Body 是一个 buffer 对象，只能读取 1 次
			requestBody, _ := ioutil.ReadAll(ctx.Request.Body)
			// 缓存拷贝 1 份，再还给他，以供后续的操作
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// req 开始
		start := time.Now()

		ctx.Next()

		cost := time.Since(start)
		respStatus := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", respStatus),
			zap.String("request", fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.String())),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", fmt.Sprintf("%.3f ms", float64(cost.Nanoseconds())/1e6)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			// 请求 body
			logFields = append(logFields,zap.String("Request body", string(requestBody)))
			// 响应 body
			logFields = append(logFields, zap.String("Response body", w.body.String()))
		} 

		if respStatus > 400 && respStatus <= 499 {
			// 除了 400 StatusBadRequest 以外，warn 提示一下，常见的 404 开发时都要注意

			zap.L().Warn(
				fmt.Sprintf("HTTP 警告 %v",respStatus),
				logFields...
			)

		} else if respStatus >= 500 && respStatus <= 599 {
			// 除了内部错误，记录 error
			
			zap.L().Warn(
				fmt.Sprintf("HTTP 错误 %v",respStatus),
				logFields...
			)
		} else {
			zap.L().Debug("HTTP 访问日志", logFields...)
		}
	}
}



// 缓存拷贝一份 "返回响应" （有些数据没啥接口访问）
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
