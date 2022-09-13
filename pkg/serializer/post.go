package serializer

import (
	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
)

// Post 帖子序列化器
type Post struct {
	PostID        int64  `json:"post_id"` 
	Title   	  string `json:"title"`
	Content       string `json:"content"`
	AuthorID      int64  `json:"author_id"`
	CommunityID   string   `json:"community_id"` 
	Status        int32  `json:"-"`
	UpdatedAt     int64  `json:"updateed_at"`
}


// BuildPost 序列化帖子
func BuildPost(item models.Post) Post {
	community, _ := mysql.GetCommunityDatailByID(item.CommunityID)

	return Post{
		UpdatedAt: item.UpdatedAt.Unix(),
		Content: item.Content,
		Title: item.Title,
		AuthorID: item.AuthorID,
		CommunityID: community.CommunityName,
	}
}

