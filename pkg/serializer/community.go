package serializer

import "github.com/sjxiang/bluebell/models"


// Community 社区话题序列化器
type Community struct {
	ID            uint64  `json:"id"`
	CommunityID   int64   `json:"community_id"`
	CommunityName string  `json:"community_name"`
	Introduction  string  `json:"introduction,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
}


// BuildCommunity 序列化视频
func BuildCommunity(item models.Community) Community {
	return Community{
		CommunityID: item.CommunityID,
		CommunityName: item.CommunityName,

		// CreatedAt: item.CreatedAt.Unix(),  // unix 时间戳

		// Avatar:     item.AvatarURL(),  // 签名的 key 
		// View:       item.View(),
	}
}


// BuildCommunity 序列化视频列表
func BuildCommunityList(items []models.Community) (communityList []Community) {

	for _, item := range items {

		community := BuildCommunity(item)
		communityList = append(communityList, community)
	}

	return communityList
}


func BuildCommunityDetail(item models.Community) Community {
	return Community{
		CommunityID: item.CommunityID,
		CommunityName: item.CommunityName,
		Introduction: item.Introduction,
		CreatedAt: item.CreatedAt.Unix(),  // unix 时间戳
	}
}