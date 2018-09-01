package db

import "github.com/jinzhu/gorm"

type Comic struct {
	ID      int
	Title   string `gorm:"size:255"`
	Summary string `sql:"type:text"`
}

type Author struct {
	ID   int
	Name string `gorm:"size:255"`
	Info string `sql:"type:text"`
}

type ComicData struct {
	ID       int
	ComicID  int
	Title    string `gorm:"size:255"`
	Episode  int
	Page     int
	Comments []*Comment
}

type ComicAuthor struct {
	ComicID  int
	AuthorID int
}

func (ca ComicAuthor) TableName() string {
	return "comic_author"
}

type User struct {
	gorm.Model
	Username   string
	Credential UserCredential
}

type UserCredential struct {
	gorm.Model
	UserID   int
	Password string
}

type Comment struct {
	gorm.Model
	ComicID  int
	UserID   int
	Content  string `sql:"type:text"`
	AssetUrl string
}

type Friendship struct {
	gorm.Model
	UserID       int
	FollowUserID int
}
