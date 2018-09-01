package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nagisa-intern/ebutuoy/server/db"
	"golang.org/x/crypto/bcrypt"
)

func (r *Router) PostMe(c echo.Context) error {
	req := &struct {
		Username string
		Password string
	}{}
	c.Bind(req)
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, M{"username or password is empty"})
	}
	// TODO: 既存チェック

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, M{"something wrong in password hashing"})
	}

	user := &db.User{}
	user.Username = req.Username
	user.Credential.Password = string(hashedPassword)
	r.db.Create(user)

	fmt.Println(*user)
	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, M{"something wrong in getting session"})
	}
	sess.Values["user_id"] = user.ID
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusCreated)
}

func (r *Router) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user := &db.User{}

	r.db.First(user, userID)
	if user.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (r *Router) GetMe(c echo.Context) error {
	userID := c.Get("userID")
	user := &db.User{}

	r.db.First(user, userID)
	if user.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (r *Router) PostFriendship(c echo.Context) error {
	myID := c.Get("userID").(int)
	followUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"id is not a number"})
	}

	friendship := &db.Friendship{}

	r.db.Where("user_id = ? AND follow_user_id = ?", myID, followUserID).First(friendship)
	if friendship.ID != 0 {
		return c.JSON(http.StatusConflict, M{"you already follow user"})
	}

	friendship = &db.Friendship{
		UserID:       myID,
		FollowUserID: followUserID,
	}
	r.db.Create(friendship)

	return c.NoContent(http.StatusCreated)
}

func (r *Router) DeleteFriendship(c echo.Context) error {
	myID := c.Get("userID").(int)
	followUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"id is not a number"})
	}

	friendship := &db.Friendship{}

	r.db.Where("user_id = ? AND follow_user_id = ?", myID, followUserID).First(friendship)
	if friendship.ID == 0 {
		return c.JSON(http.StatusNotFound, M{"friendship is not found"})
	}

	r.db.Delete(friendship)
	return c.NoContent(http.StatusOK)
}
