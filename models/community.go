package models





// 论坛
type Community struct {
	Base 
	
	CommunityID   int64  `gorm:"type:varchar(64); column:community_id;   not null"`
	CommunityName string `gorm:"type:varchar(128);column:community_name; not null; UNIQUE"`
	Introduction  string `gorm:"type:varchar(256)"`

	CommonTimestampsField
}

