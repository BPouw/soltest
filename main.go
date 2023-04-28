package main

import (
	"webshop/api/config"
	"webshop/api/route"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//run database
	config.ConnectDB()

	// routes
	route.UserRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}
