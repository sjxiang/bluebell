package serializer

import "github.com/sjxiang/bluebell/models"


// Community 社区话题序列化器
type Community struct {
	ID            uint64  `json:"id"`
	CommunityID   int64   `json:"community_id"`
	CommunityName string  `json:"community_name"`
	Introduction  string  `json:"introduction"`
	CreatedAt     int64   `json:"created_at"`
}


// BuildCommunity 序列化视频
func BuildCommunity(item models.Community) Community {
	return Community{
		ID: item.ID,
		CommunityID: item.CommunityID,
		CommunityName: item.CommunityName,
		Introduction: item.Introduction,
		CreatedAt: item.CreatedAt.Unix(),  // unix 时间戳
		
		// Avatar:     item.AvatarURL(),  // 签名的 key 
		// View:       item.View(),
	}
}


// BuildCommunity 序列化视频列表
func BuildCommunitys(items []models.Community) (communitys []Community) {

	for _, item := range items {

		community := BuildCommunity(item)
		communitys = append(communitys, community)
	}

	return communitys
}
