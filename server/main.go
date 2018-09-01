package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nagisa-intern/ebutuoy/server/db"
)

func main() {
	dbManager := db.New()
	if err := dbManager.Connect(); err != nil {
		panic(err)
	}

	dbManager.Sync()

	e := echo.New()

	e.GET("/ping", pong)
	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
