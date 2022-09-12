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


func GetCommunityDatailByID(id int64) (models.Community, error) {
	var community models.Community

	result := DB.Where("community_id  = ?", id).Find(&community)
	
	if result.Error != nil {
		return community, result.Error  // 数据库查询出错
	}

	if community.ID > 0 {
		return community, nil
	}

	return community, errors.New("查询为空")
}