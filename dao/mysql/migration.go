package mysql

import "github.com/sjxiang/bluebell/models"


func migration() {
	_ = DB.AutoMigrate(&models.User{})
	_ = DB.AutoMigrate(&models.Community{})
	_ = DB.AutoMigrate(&models.Post{})
}
