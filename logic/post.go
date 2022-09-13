package logic

import (
	"errors"
	"strconv"

	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
	"github.com/sjxiang/bluebell/pkg/snowflake"
	"github.com/sjxiang/bluebell/requests"
)


func CreatePost(p *requests.ParamCreatePost, authorID int64) error {

	// 生成 id
	postID := snowflake.GetID()

	communityid, _ := strconv.ParseInt(p.CommunityID, 10, 64)

	//  保存到数据库
	post := models.Post{
		PostID: postID,
		Title: p.Title,
		Content: p.Content,
		CommunityID: communityid,
		AuthorID: authorID,
	}

	if ok := mysql.CreatePost(&post); !ok {
		return errors.New("创建 post 失败")
	}

	return nil
}


func GetPostByID(pid int64) (*models.Post, error) {
	return mysql.GetPostByID(pid)
}