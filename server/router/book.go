package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nagisa-intern/ebutuoy/server/db"
)

type Comic struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Authors []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info string `json:"info"`
	} `json:"authors"`
}

func (r *Router) GetComics(c echo.Context) error {
	comics := []*Comic{}
	r.db.Find(&comics)
	for _, c := range comics {
		r.db.Table("authors").Joins("LEFT JOIN comic_author ON comic_author.author_id = authors.id").Where("comic_author.comic_id = ?", c.ID).Find(&c.Authors)
		fmt.Println(c.Authors)
	}
	return c.JSON(http.StatusOK, comics)
}

func (r *Router) GetComicByID(c echo.Context) error {
	comicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"comic id is not number"})
	}

	res := &struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Summary string `json:"summary"`
		Authors []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Info string `json:"info"`
		} `json:"authors"`
		Episodes []struct {
			Title   string `json:"title"`
			Episode int    `json:"episode"`
			Page    int    `json:"page"`
		} `json:"episodes"`
	}{}
	r.db.Table("comics").First(res, comicID)
	r.db.Table("authors").Joins("LEFT JOIN comic_author ON comic_author.author_id = authors.id").Where("comic_author.comic_id = ?", res.ID).Find(&res.Authors)
	r.db.Table("comic_data").Where("comic_id = ?", comicID).Find(&res.Episodes)
	return c.JSON(http.StatusOK, res)
}

func (r *Router) GetCommentsByID(c echo.Context) error {
	comicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"comic id is not number"})
	}

	comments := &[]db.Comment{}
	r.db.Where("comic_id = ?", comicID).Find(comments)
	return c.JSON(http.StatusOK, comments)
}

func (r *Router) PostCommentsByID(c echo.Context) error {
	userID := c.Get("userID").(int)
	comicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"comic id is not number"})
	}

	req := &struct {
		Episode  int    `json:"episode"`
		Content  string `json:"content"`
		AssetUrl string `json:"asset_url"`
	}{}
	c.Bind(req)

	comment := &db.Comment{
		ComicID:  comicID,
		UserID:   userID,
		Content:  req.Content,
		AssetUrl: req.AssetUrl,
	}
	r.db.Create(comment)
	return c.NoContent(http.StatusCreated)
}
