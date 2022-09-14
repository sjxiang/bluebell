package logic

import (
	"strconv"

	"github.com/sjxiang/bluebell/dao/redis"
	"github.com/sjxiang/bluebell/requests"
)

/*

投票的几种情况

*/


func PostVote(userID int64, p *requests.ParamPostVote) error {
	userid := strconv.Itoa(int(userID))
	postid := strconv.Itoa(int(p.PostID))

	return redis.VoteForPost(userid, postid, p.Direction)
}
