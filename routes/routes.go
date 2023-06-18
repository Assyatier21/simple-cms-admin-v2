package routes

import (
	"github.com/assyatier21/simple-cms-admin-v2/internal/delivery/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/assyatier21/simple-cms-admin-v2/docs"
)

func InitRoutes(handler api.DeliveryHandler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.CORS())
}
