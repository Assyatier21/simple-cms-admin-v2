package postgres

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
)

func (r *repository) buildArticleResponse(scenario string, req interface{}) (article entity.ArticleResponse) {
	switch scenario {
	case "insert":
		if insertReq, ok := req.(*entity.InsertArticleRequest); ok {
			article = entity.ArticleResponse{
				ID:           insertReq.ID,
				Title:        insertReq.Title,
				Slug:         insertReq.Slug,
				HTMLContent:  insertReq.HTMLContent,
				Metadata:     json.RawMessage(insertReq.Metadata),
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				CategoryList: r.buildCategoryList(insertReq.CategoryIDs),
			}
		}
	case "update":
		if updateReq, ok := req.(*entity.UpdateArticleRequest); ok {
			article = entity.ArticleResponse{
				ID:           updateReq.ID,
				Title:        updateReq.Title,
				Slug:         updateReq.Slug,
				HTMLContent:  updateReq.HTMLContent,
				Metadata:     json.RawMessage(updateReq.Metadata),
				CreatedAt:    article.CreatedAt,
				UpdatedAt:    time.Now().Format(time.RFC3339),
				CategoryList: r.buildCategoryList(updateReq.CategoryIDs),
			}
		}
	}

	return article
}

func (r *repository) buildCategoryList(categoryIDs []int) []entity.CategoryResponse {
	categories, err := r.GetCategoriesByIDs(context.Background(), categoryIDs)
	if err != nil {
		log.Println("[Repository][Postgres][buildCategoryList] error failed to get category by ids, err: ", err)
		return nil
	}

	return categories
}
