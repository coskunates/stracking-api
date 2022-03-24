package app

import (
	"github.com/labstack/echo/v4"
	"stock/controllers"
)

func handleRoutes(e *echo.Echo) {
	e.GET("/ping", controllers.PingController.Ping)
	e.POST("/positions", controllers.PositionController.Create)
}
