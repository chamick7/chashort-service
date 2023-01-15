package controller

import (
	"context"

	"github.com/chamick7/short-service/module/url"
	"github.com/chamick7/short-service/sqlc"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, ec *echo.Echo, conn *sqlc.Queries) {
	urlService := url.New(ctx, conn)

	redirectGroup := ec.Group("")
	redirectRouter := RedirectRouter{group: redirectGroup, urlService: urlService}
	redirectRouter.register()

	shortGroup := ec.Group("/short")
	shortRouter := ShortRouter{group: shortGroup, urlService: urlService}
	shortRouter.register()
}
