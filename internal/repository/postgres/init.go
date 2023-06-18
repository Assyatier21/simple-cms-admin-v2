package postgres

import (
	"context"
	"database/sql"

	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
)

type RepositoryHandler interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]entity.ArticleResponse, error)
	GetArticleDetails(ctx context.Context, id string) (entity.ArticleResponse, error)
	InsertArticle(ctx context.Context, article entity.Article) (entity.ArticleResponse, error)
	UpdateArticle(ctx context.Context, article entity.Article) (entity.ArticleResponse, error)
	DeleteArticle(ctx context.Context, id string) error

	GetCategoryTree(ctx context.Context) ([]entity.Category, error)
	GetCategoryDetails(ctx context.Context, id int) (entity.Category, error)
	InsertCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
