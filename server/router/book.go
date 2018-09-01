package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Comic struct {
	ID      int
	Title   string
	Summary string
	Authors []struct {
		ID   int
		Name string
		Info string
	}
}

func (r *Router) GetComics(c echo.Context) error {
	comics := []*Comic{}
	r.db.Find(&comics)
	fmt.Println(comics[0])
	for _, c := range comics {
		r.db.Table("authors").Joins("LEFT JOIN comic_author ON comic_author.author_id = authors.id").Where("comic_author.comic_id = ?", c.ID).Find(&c.Authors)
		fmt.Println(c.Authors)
	}
	return c.JSON(http.StatusOK, comics)
}
