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



func GetPostList(pageSize, pageNum int64) ([]models.Post, error) {


	// 判断是否需要分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1  // 默认 等于 *
	}


	// 查找

	var postList []models.Post

	DB.Table("post").Select("title", "community_id").Limit(int(pageSize)).Offset(int(offsetVal)).Find(&postList)
	if len(postList) > 0 {
		return postList, nil
	}

	return nil, errors.New("查无此贴")
}