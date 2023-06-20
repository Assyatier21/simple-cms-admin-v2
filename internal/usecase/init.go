package usecase

import (
	"context"

	elastic "github.com/assyatier21/simple-cms-admin-v2/internal/repository/elasticsearch"
	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres"
	"github.com/assyatier21/simple-cms-admin-v2/models"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
)

type UsecaseHandler interface {
	GetArticles(ctx context.Context, req entity.GetArticlesRequest) models.StandardResponseReq
	GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) models.StandardResponseReq
	InsertArticle(ctx context.Context, req entity.InsertArticleRequest) models.StandardResponseReq
	UpdateArticle(ctx context.Context, req entity.UpdateArticleRequest) models.StandardResponseReq
	DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) models.StandardResponseReq
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
