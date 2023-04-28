package route

import (
	"webshop/api/controller"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	//All routes related to users comes here
	e.POST("/item", controller.CreateItem)
	e.GET("/item/:itemId", controller.GetItem)
}
