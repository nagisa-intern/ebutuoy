package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nagisa-intern/ebutuoy/server/db"
)

func (r *Router) GetTimeline(c echo.Context) error {
	myID := c.Get("userID").(int)
	comments := &[]db.Comment{}
	r.db.Joins("LEFT JOIN friendships ON friendships.follow_user_id = comments.user_id").Where("friendships.user_id = ? OR comments.user_id = ?", myID, myID).Order("created_at desc").Limit(10).Find(comments)
	return c.JSON(http.StatusOK, comments)
}
