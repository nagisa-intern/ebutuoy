package router

import (
	"github.com/jinzhu/gorm"
	"github.com/nagisa-intern/ebutuoy/server/db"
)

type Router struct {
	db *gorm.DB
}

type M struct {
	Message string `json:"message"`
}

func New() (*Router, error) {
	r := &Router{}
	_db, err := db.New()
	if err != nil {
		return nil, err
	}
	r.db = _db

	return r, nil
}
