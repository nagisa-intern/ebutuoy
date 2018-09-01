package router

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nagisa-intern/ebutuoy/server/db"
	"github.com/srinathgs/mysqlstore"
	"golang.org/x/crypto/bcrypt"
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

		return next(c)
	}
}

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
