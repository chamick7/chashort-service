package controller

import (
	"github.com/chamick7/short-service/module/url"
	"github.com/labstack/echo/v4"
)

type RedirectRouter struct {
	group      *echo.Group
	urlService url.UrlService
}

func (r *RedirectRouter) register() {
	r.group.GET("/:shortId", r.urlService.RedirectTO)
}
