package main

import (
	"github.com/assyatier21/simple-cms-admin-v2/config"
	"github.com/assyatier21/simple-cms-admin-v2/driver"
	"github.com/assyatier21/simple-cms-admin-v2/internal/delivery/api"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/elasticsearch"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres"
	"github.com/assyatier21/simple-cms-admin-v2/internal/usecase"
	"github.com/assyatier21/simple-cms-admin-v2/routes"
	"github.com/assyatier21/simple-cms-admin-v2/utils/helper"
)

// @title           Swagger Simple CMS Admin
// @version         2.0
// @description     This is a documentation of Simple Content Management System V2.
func main() {
	// Load Config
	cfg := config.Load()

	dbClient := driver.InitPostgres(cfg.PostgresConfig)
	esClient := driver.InitElasticClient(cfg.ElasticConfig)

	postgresRepository := postgres.NewRepository(dbClient)
	elasticRepository := elasticsearch.NewElasticRepository(esClient, cfg.ElasticConfig)

	usecase := usecase.NewUsecase(postgresRepository, elasticRepository)
	delivery := api.NewHandler(usecase)

	echo := routes.InitRoutes(delivery)
	helper.UseCustomValidatorHandler(echo)

	echo.Start(":8800")
}
