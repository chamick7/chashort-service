package controller

import (
	"github.com/chamick7/short-service/module/url"
	"github.com/labstack/echo/v4"
)

type ShortRouter struct {
	group      *echo.Group
	urlService url.UrlService
}

func (r *ShortRouter) register() {
	r.group.POST("", r.urlService.ShortURL)
}
