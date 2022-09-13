package mysql

import (
	"errors"

	"github.com/sjxiang/bluebell/models"
)


func CreatePost(post *models.Post) bool {
	
	
	// 保存到数据库
	DB.Create(post)
	return post.ID > 0
}


func GetPostByID(pid int64) (*models.Post, error) {
	var post models.Post
	DB.Where("post_id = ?", pid).Find(&post)
	if post.ID > 0 {
		return &post, nil
	}

	return nil, errors.New("查无此文")
}