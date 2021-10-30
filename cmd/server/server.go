package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickeasters/nobones-api/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handlers.GetBonesJSON)
	e.GET("/bones", handlers.GetBones)
	e.POST("/admission", handlers.AdmissionWebhook)
	e.Logger.Fatal(e.Start(":3000"))
}
