package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nagisa-intern/ebutuoy/server/db"
	"github.com/srinathgs/mysqlstore"
	"golang.org/x/crypto/bcrypt"
)

func (r *Router) CreateSessionStore() (*mysqlstore.MySQLStore, error) {
	return mysqlstore.NewMySQLStoreFromConnection(r.db.DB(), "sessions", "/", 60*60*24*14, []byte("popo"))
}

func (r *Router) WithLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, M{"something wrong in getting session"})
		}

		if sess.Values["user_id"] == nil {
			return c.JSON(http.StatusForbidden, M{"please login"})
		}
		c.Set("userID", int(sess.Values["user_id"].(uint)))

		return next(c)
	}
}

func (r *Router) PostLogin(c echo.Context) error {
	req := &struct {
		Username string
		Password string
	}{}

	c.Bind(req)

	user := &db.User{
		Username: req.Username,
	}

	credential := &db.UserCredential{}

	r.db.First(user)
	r.db.Model(user).Related(credential)

	err := bcrypt.CompareHashAndPassword([]byte(credential.Password), []byte(req.Password))

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusForbidden, M{"username or password is incorrect"})
	}

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, M{"something wrong in getting session"})
	}
	sess.Values["user_id"] = user.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		fmt.Println(err)
	}

	return c.NoContent(http.StatusOK)
}

func (r *Router) PostLogout(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, M{"something wrong in getting session"})
	}
	sess.Values["user_id"] = nil
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		fmt.Println(err)
	}

	return c.NoContent(http.StatusOK)
}
