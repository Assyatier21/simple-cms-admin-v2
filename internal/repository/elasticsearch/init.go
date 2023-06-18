package elasticsearch

import (
	"context"

	"github.com/assyatier21/simple-cms-admin-v2/config"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	"github.com/olivere/elastic/v7"
)

type ElasticHandler interface {
	GetArticles(ctx context.Context, limit int, offset int, sort_by string, order_by bool) ([]entity.ArticleResponse, error)
	GetArticleDetails(ctx context.Context, query elastic.Query) (entity.ArticleResponse, error)
	InsertArticle(ctx context.Context, article entity.ArticleResponse) error
	UpdateArticle(ctx context.Context, article entity.ArticleResponse) error
	DeleteArticle(ctx context.Context, id string) error

	GetCategoryTree(ctx context.Context, limit int, offset int) ([]entity.Category, error)
	GetCategoryDetails(ctx context.Context, query elastic.Query) (entity.Category, error)
	InsertCategory(ctx context.Context, category entity.Category) error
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type elasticRepository struct {
	cfg config.ElasticConfig
	es  *elastic.Client
}

func NewElasticRepository(es *elastic.Client, cfg config.ElasticConfig) ElasticHandler {
	return &elasticRepository{
		cfg: cfg,
		es:  es,
	}
}
