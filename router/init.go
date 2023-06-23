package router

import (
	"github.com/assyatier21/simple-cms-admin-v2/internal/handler/api"
	_ "github.com/assyatier21/simple-cms-admin-v2/middleware"
	"github.com/labstack/echo/v4"

	_ "github.com/assyatier21/simple-cms-admin-v2/docs"
)

func InitRouter(server *echo.Echo, handler api.DeliveryHandler) {
	InitArticleRouter(server, handler)
	InitCategoryRouter(server, handler)
}
