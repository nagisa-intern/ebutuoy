package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nagisa-intern/ebutuoy/server/router"
)

func main() {

	e := echo.New()
	e.GET("/ping", pong)

	r, err := router.New()
	if err != nil {
		panic(err)
	}
	store, err := r.CreateSessionStore()
	if err != nil {
		panic(err)
	}
	e.Use(session.Middleware(store))

	api := e.Group("/api")

	api.POST("/me", r.PostMe)
	api.POST("/login", r.PostLogin)

	api.Use(r.WithLogin)
	api.GET("/logined_ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
