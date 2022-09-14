package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	oneWeekTime = 7 * 24 * 3600
)

func VoteForPost(userID , postID string, direction int64) error {
	// 1. 判断投票限制
	// 去 redis 取帖子发布时间
	postTime := Redis.Client.ZScore(Redis.Context, getRedisKey(KeyPostTimeZset), postID).Val()

	if time.Now().Unix() - int64(postTime) > oneWeekTime {
		return errors.New("投票时间已过")
	}
	
	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子的投票纪录
	ov := Redis.Client.ZScore(
		Redis.Context, 
		getRedisKey(KeyPostVotedZsetPrefix + postID), 
		userID).Val()
	
	if ov > float64(direction) {
		direction = 1
	} else {
		direction = -1
	}

	diff := math.Abs(ov - float64(direction))  // 计算两次差值
	_, err := Redis.Client.ZIncrBy(Redis.Context, 
		getRedisKey(KeyPostScoreZset), float64(direction) * diff, postID).Result()

	if err != nil {
		return err
	}

	// 3. 记录用户为该帖子投票的数据
	if direction == 0 {
		_, err := Redis.Client.ZRem(
			Redis.Context, 
			getRedisKey(KeyPostVotedZsetPrefix + postID), 
			postID).Result()
		return err
	} else {
		_, err := Redis.Client.ZAdd(
			Redis.Context, 
			getRedisKey(KeyPostVotedZsetPrefix + postID), 
			&redis.Z{Score: float64(direction), Member: userID}).Result()
		return err
	}
}