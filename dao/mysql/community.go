package mysql

import (
	"errors"

	"github.com/sjxiang/bluebell/models"
)


func GetCommunityList() ([]*models.Community, error) {
	
	var communityList []*models.Community

	DB.Select("community_id", "community_name").Find(&communityList)

	if len(communityList) == 0 {
		return nil, errors.New("数据库里找不到")
	}

	return communityList, nil
}