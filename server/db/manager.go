package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() *DBManager {
	return &DBManager{}
}

type DBManager struct {
	DB *gorm.DB
}

func (d *DBManager) Connect() error {
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		username = "root"
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "mysql"
	}

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "43306"
	}

	dbname := os.Getenv("MYSQL_DB_NAME")
	if dbname == "" {
		dbname = "jihen_intern"
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname))
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}

func (d *DBManager) Sync() {
	d.DB.AutoMigrate(&Comic{}, &Author{}, &ComicData{}, &ComicAuthor{}, &User{}, &UserCredential{}, &Comment{}, &Friendship{})
}
