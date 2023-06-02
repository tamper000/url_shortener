package main

import (
	"url_shortener/database"
	"url_shortener/routing"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Init()
	http_server()
}

func http_server() {
	e := echo.New()

	e.File("/", "web/static/index.html")
	e.Static("/assets", "web/assets")

	e.POST("/api/create", routing.Create)
	e.GET("/:shortID", routing.GetShorten)

	e.Logger.Fatal(e.Start(":8080"))
}
