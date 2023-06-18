package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	query "github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres/queries"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	msg "github.com/assyatier21/simple-cms-admin-v2/models/lib"
)

func (r *repository) GetArticles(ctx context.Context, limit int, offset int) ([]entity.ArticleResponse, error) {
	var (
		articles []entity.ArticleResponse
		rows     *sql.Rows
		err      error
	)
	rows, err = r.db.Query(query.GET_ARTICLES, limit, offset)
	if err != nil {
		log.Println("[Repository][GetArticles] failed to get list of articles, err: ", err)
		return nil, err
	}

	for rows.Next() {
		var (
			temp         entity.ArticleResponse
			byteMetadata []byte
		)

		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetArticles] failed to scan article, err :", err)
			return nil, err
		}

		err = r.db.QueryRow(query.GET_METADATA, temp.Id).Scan(&byteMetadata)
		if err != nil {
			log.Println("[Repository][GetArticles] failed to scan metadata, err :", err)
			return nil, err
		}

		json.Unmarshal(byteMetadata, &temp.MetaData)
		articles = append(articles, temp)
	}

	if len(articles) == 0 {
		return []entity.ArticleResponse{}, nil
	}

	return articles, nil
}

func (r *repository) GetArticleDetails(ctx context.Context, id string) (entity.ArticleResponse, error) {
	var (
		article      entity.ArticleResponse
		err          error
		byteMetadata []byte
	)

	err = r.db.QueryRow(query.GET_ARTICLE_DETAILS, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	err = r.db.QueryRow(query.GET_METADATA, article.Id).Scan(&byteMetadata)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan metadata, err :", err)
		return entity.ArticleResponse{}, err
	}
	json.Unmarshal(byteMetadata, &article.MetaData)

	return article, nil
}
func (r *repository) InsertArticle(ctx context.Context, article entity.Article) (entity.ArticleResponse, error) {
	var (
		ArticleReponse entity.ArticleResponse
		err            error
	)

	marshaled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to insert article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	_, err = r.db.Exec(query.INSERT_ARTICLE, article.Id, article.Title, article.Slug, article.HtmlContent, article.CategoryID, marshaled_metadata, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to insert article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	ArticleReponse, err = r.GetArticleDetails(context.Background(), article.Id)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to get article details response, err: ", err)
		return entity.ArticleResponse{}, err
	}

	return ArticleReponse, nil
}
func (r *repository) UpdateArticle(ctx context.Context, article entity.Article) (entity.ArticleResponse, error) {
	var (
		ArticleReponse entity.ArticleResponse
		rows           sql.Result
		err            error
	)

	marshaled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to update article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	rows, err = r.db.Exec(query.UPDATE_ARTICLE, &article.Title, &article.Slug, &article.HtmlContent, &article.CategoryID, marshaled_metadata, &article.UpdatedAt, &article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to update article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	ArticleReponse, err = r.GetArticleDetails(context.Background(), article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to get article details response, err: ", err)
		return entity.ArticleResponse{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return entity.ArticleResponse{}, nil
	}

	return ArticleReponse, nil
}
func (r *repository) DeleteArticle(ctx context.Context, id string) error {
	rows, err := r.db.Exec(query.DELETE_ARTICLE, id)
	if err != nil {
		log.Println("[Repository][DeleteArticle] failed to delete article, err: ", err)
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return msg.ERROR_NO_ROWS_AFFECTED
	}

	return nil
}
