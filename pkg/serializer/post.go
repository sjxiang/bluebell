package serializer

import (
	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
)

// Post 帖子序列化器
type Post struct {
	PostID        int64  `json:"post_id,omitempty"` 
	Title   	  string `json:"title"`
	Content       string `json:"content,omitempty"`
	AuthorID      int64  `json:"author_id,omitempty"`
	CommunityID   string `json:"community_id,omitempty"` 
	Status        int32  `json:"-"`
	UpdatedAt     int64  `json:"updateed_at,omitempty"`
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



func BuildPostList(items []models.Post) (postList []Post) {

	for _, item := range items {

		post := BuildPost(item)
		postList = append(postList, post)
	}

	return postList
}