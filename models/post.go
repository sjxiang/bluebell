package models



type Post struct {
	Base 
	
	PostID        int64  `gorm:"type:varchar(64);not null"`
	Title   	  string `gorm:"type:varchar(128);not null;UNIQUE"`
	Content       string `gorm:"type:varchar(1024)"`
	AuthorID      int64  `gorm:"type:varchar(64);UNIQUE"`
	CommunityID   int64  `gorm:"type:varchar(128);not null"`
	Status        int32  `gorm:"not null;default:1"`

	CommonTimestampsField
}

