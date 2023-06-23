package main

import (
	"fmt"

	"github.com/assyatier21/simple-cms-admin-v2/config"
	"github.com/assyatier21/simple-cms-admin-v2/driver"
	"github.com/assyatier21/simple-cms-admin-v2/internal/handler/api"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/elasticsearch"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres"
	"github.com/assyatier21/simple-cms-admin-v2/internal/usecase"
	"github.com/assyatier21/simple-cms-admin-v2/middleware"
	"github.com/assyatier21/simple-cms-admin-v2/router"
	"github.com/labstack/echo/v4"
)

// @title           Swagger Simple CMS Admin
// @version         2.0
// @description     This is a documentation of Simple Content Management System V2.
func main() {
	server := echo.New()

	// Load Config
	cfg := config.Load()

	dbClient := driver.InitPostgres(cfg.PostgresConfig)
	esClient := driver.InitElasticClient(cfg.ElasticConfig)

	postgresRepository := postgres.NewRepository(dbClient)
	elasticRepository := elasticsearch.NewElasticRepository(esClient, cfg.ElasticConfig)

	usecase := usecase.NewUsecase(postgresRepository, elasticRepository)
	handler := api.NewHandler(usecase)

	router.InitRouter(server, handler)
	middleware.InitMiddlewares(server)

	host := fmt.Sprintf("%s:%s", cfg.ApplicationConfig.Host, cfg.ApplicationConfig.Port)
	server.Start(host)
}
