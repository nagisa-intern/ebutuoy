package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nagisa-intern/ebutuoy/server/router"
)

func main() {

	e := echo.New()

	r, err := router.New()
	if err != nil {
		panic(err)
	}
	store, err := r.CreateSessionStore()
	if err != nil {
		panic(err)
	}
	e.Use(session.Middleware(store))

	e.GET("/ping", pong)
	e.POST("/me", r.PostMe)
	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
