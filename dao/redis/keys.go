package redis



// Redis Key

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZset        = "post:time"   // 帖子及发帖时间
	KeyPostScoreZset       = "post:score"  // 帖子及投票的分数
	KeyPostVotedZsetPrefix = "post:voted:"          // 记录用户及投票类型
)


// 是 Redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}