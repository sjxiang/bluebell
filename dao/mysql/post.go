package mysql

import "github.com/sjxiang/bluebell/models"


func CreatePost(post *models.Post) bool {
	
	
	// 保存到数据库
	DB.Create(post)
	return post.ID > 0
}