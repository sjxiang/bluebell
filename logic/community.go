package logic

import (
	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
)



func GetCommunityList() ([]models.Community, error) {
	// 查找所有的 community 并返回
	return mysql.GetCommunityList()
}