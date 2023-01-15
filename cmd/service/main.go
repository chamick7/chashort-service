package main

import (
	"context"
	"net/http"

	"github.com/chamick7/short-service/controller"
	"github.com/chamick7/short-service/database"
	"github.com/chamick7/short-service/utils/validate"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	ctx := context.Background()

	conn := database.Init()

	e := echo.New()
	validate.Init(e)

	controller.Register(ctx, e, conn)

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":5000"))
}
