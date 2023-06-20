package postgres

import (
	"context"
	"encoding/json"
	"log"

	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres/queries"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	"github.com/assyatier21/simple-cms-admin-v2/models/lib"
	"github.com/lib/pq"
)

func (r *repository) GetArticles(ctx context.Context, req entity.GetArticlesRequest) ([]entity.ArticleResponse, error) {
	var (
		articles = []entity.ArticleResponse{}
	)

	rows, err := r.db.Query(queries.GET_ARTICLES, req.Limit, req.Offset)
	if err != nil {
		log.Println("[Repository][Postgres][GetArticles] error failed to exec query, err: ", err)
		return articles, nil

	}
	defer rows.Close()

	for rows.Next() {
		var article entity.ArticleResponse
		var categoryJSON json.RawMessage

		err := rows.Scan(&article.ID, &article.Title, &article.Slug, &article.HTMLContent, &article.Metadata, &article.CreatedAt, &article.UpdatedAt, &categoryJSON)
		if err != nil {
			log.Println("[Repository][Postgres][GetArticles] error failed to scan data, err: ", err)
			return articles, nil
		}

		err = json.Unmarshal(categoryJSON, &article.CategoryList)
		if err != nil {
			log.Println("[Repository][Postgres][GetArticles] error failed to unmarshal categories, err: ", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *repository) GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) (entity.ArticleResponse, error) {
	var (
		article = entity.ArticleResponse{}
	)
	rows, err := r.db.Query(queries.GET_ARTICLE_DETAILS, req.ID)
	if err != nil {
		log.Println("[Repository][Postgres][GetArticleDetails] error failed to exec query, err: ", err)
		return article, err
	}
	defer rows.Close()

	for rows.Next() {
		var categories []entity.CategoryResponse
		err := rows.Scan(
			&article.ID, &article.Title, &article.Slug, &article.HTMLContent,
			&article.Metadata, &article.CreatedAt, &article.UpdatedAt,
			pq.Array(&categories),
		)
		if err != nil {
			log.Println("[Repository][Postgres][GetArticleDetails] error failed to scan rows, err: ", err)
			return article, err
		}
		article.CategoryList = categories
	}

	return article, err
}

func (r *repository) InsertArticle(ctx context.Context, article entity.InsertArticleRequest) (entity.ArticleResponse, error) {
	_, err := r.db.Exec(queries.INSERT_ARTICLE, article.ID, article.Title, article.Slug, article.HTMLContent, article.CategoryIDs, article.Metadata, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][Postgres][InsertArticle] error failed to insert article, err: ", err)
		return entity.ArticleResponse{}, err
	}
	articleResponse := r.buildArticleResponse("insert", article)

	return articleResponse, nil
}

func (r *repository) UpdateArticle(ctx context.Context, article entity.UpdateArticleRequest) (entity.ArticleResponse, error) {
	result, err := r.db.Exec(queries.UPDATE_ARTICLE, article.Title, article.Slug, article.HTMLContent, article.CategoryIDs, article.Metadata, article.CreatedAt, article.UpdatedAt, article.ID)
	if err != nil {
		log.Println("[Repository][Postgres][UpdateArticle] error failed to update article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.ArticleResponse{}, lib.ErrorNoRowsAffected
	}

	articleResponse := r.buildArticleResponse("update", article)

	return articleResponse, nil
}

func (r *repository) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) error {
	result, err := r.db.Exec(queries.DELETE_ARTICLE, req.ID)
	if err != nil {
		log.Println("[Repository][Postgres][DeleteArticle] error failed to delete article, err: ", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return lib.ErrorNoRowsAffected
	}

	return nil
}
