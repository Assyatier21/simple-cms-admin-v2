package usecase

import (
	elastic "github.com/assyatier21/simple-cms-admin-v2/internal/repository/elasticsearch"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres"
)

type UsecaseHandler interface {
}

type usecase struct {
	repository postgres.RepositoryHandler
	es         elastic.ElasticHandler
}

func NewUsecase(repository postgres.RepositoryHandler, es elastic.ElasticHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
		es:         es,
	}
}
