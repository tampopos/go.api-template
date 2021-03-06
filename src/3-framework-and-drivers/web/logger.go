package web

import (
	"github.com/labstack/echo/middleware"
	entitieslog "github.com/tampopos/go.api-template/src/0-enterprise-business-rules/entities/log"
	"github.com/tampopos/go.api-template/src/3-framework-and-drivers/log"
)

func (web *web) setLogger(lvl entitieslog.Lvl) *web {
	web.echo.Logger.SetLevel(log.CnvLvl(lvl))
	web.echo.Use(middleware.Logger())
	return web
}
