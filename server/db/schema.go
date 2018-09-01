package db

import "github.com/jinzhu/gorm"

type Comic struct {
	ID      int    `json:"id"`
	Title   string `gorm:"size:255" json:"title"`
	Summary string `sql:"type:text" json:"summary"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `gorm:"size:255" json:"name"`
	Info string `sql:"type:text" json:"info"`
}

type ComicData struct {
	ID       int        `json:"id"`
	ComicID  int        `json:"comic_id"`
	Title    string     `gorm:"size:255" json:"title"`
	Episode  int        `json:"episode"`
	Page     int        `json:"page"`
	Comments []*Comment `json:"comments"`
}

type ComicAuthor struct {
	ComicID  int `json:"comic_id"`
	AuthorID int `json:"author_id"`
}

func (ca ComicAuthor) TableName() string {
	return "comic_author"
}

type User struct {
	gorm.Model
	Username   string         `json:"username"`
	Credential UserCredential `json:"-"`
}

type UserCredential struct {
	gorm.Model
	UserID   int    `json:"user_id"`
	Password string `json:"password"`
}

type Comment struct {
	gorm.Model
	ComicID  int    `json:"comic_id"`
	UserID   int    `json:"user_id"`
	Content  string `sql:"type:text" json:"content"`
	AssetUrl string `json:"asset_url"`
}

type Friendship struct {
	gorm.Model
	UserID       int `json:"user_id"`
	FollowUserID int `json:"follow_user_id"`
}
