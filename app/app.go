package app

import "github.com/labstack/echo/v4"

var (
	e *echo.Echo
)

func StartApplication() {
	e = echo.New()

	handleRoutes(e)
	handleCron()

	e.Logger.Fatal(e.Start(":8080"))
}
