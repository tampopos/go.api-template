package web

import (
	"github.com/labstack/echo"
	"github.com/tampopos/go.api-template/src/2-interface-adapters/ctrl"
)

func (web *web) setRoute() *web {
	web.echo.GET("/", func(c echo.Context) error {
		return web.container.Invoke(func(home ctrl.Home) error {
			return home.Get(c)
		})
	})
	return web
}
