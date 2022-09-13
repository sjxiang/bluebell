package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/sjxiang/bluebell/pkg/serializer"
)


// 2， 1 每隔 2 秒钟，往桶里塞 1 个。Ps. 容量就 1 个
func RateLimit(fillinterval time.Duration, cap int64) func(ctx *gin.Context) {

	bucket := ratelimit.NewBucket(fillinterval, cap)

	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {  // 每次拿走一个，看剩下多少
			// 取不到令牌，则返回响应
			ctx.JSON(http.StatusOK, serializer.Response{Msg: "限流"})
			ctx.Abort()

			return
		}

		// 拿到令牌就放行
		ctx.Next()
	}
}