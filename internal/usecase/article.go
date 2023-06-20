package usecase

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/assyatier21/simple-cms-admin-v2/models"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	"github.com/assyatier21/simple-cms-admin-v2/utils/constant"
	"github.com/assyatier21/simple-cms-admin-v2/utils/helper"
	"github.com/olivere/elastic/v7"
)

func (u *usecase) GetArticles(ctx context.Context, req entity.GetArticlesRequest) models.StandardResponseReq {
	var (
		articles = []entity.ArticleResponse{}
	)

	req.SortBy = helper.ValidateSortBy(req.SortBy)
	req.OrderByBool = helper.ValidateOrderBy(req.OrderBy)

	articles, err := u.es.GetArticles(ctx, req)
	if err != nil {
		log.Println("[Usecase][Article][Article][GetArticles] failed to get list of articles, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLES, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: articles, Error: nil}
}

func (u *usecase) GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) models.StandardResponseReq {
	var (
		article = entity.ArticleResponse{}
		query   elastic.Query
	)

	query = elastic.NewMatchQuery(constant.ID, req.ID)
	article, err := u.es.GetArticleDetails(ctx, query)
	if err != nil {
		log.Println("[Usecase][Article][Article][GetArticleDetails] failed to get article details, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLE_DETAILS, Error: err}
	}

	article = helper.FormatTimeArticleResponse(article)
	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: article, Error: nil}
}

func (u *usecase) InsertArticle(ctx context.Context, req entity.InsertArticleRequest) models.StandardResponseReq {
	reqArticle := entity.InsertArticleRequest{
		ID:          helper.GenerateUUIDString(),
		Title:       req.Title,
		Slug:        req.Slug,
		HTMLContent: req.HTMLContent,
		CategoryIDs: req.CategoryIDs,
		CreatedAt:   constant.TimeNow,
		UpdatedAt:   constant.TimeNow,
		Metadata:    req.Metadata,
	}

	articleResponse, err := u.repository.InsertArticle(ctx, reqArticle)
	if err != nil {
		log.Println("[Usecase][Article][InsertArticle] failed to insert article, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE, Error: err}
	}

	err = u.es.InsertArticle(ctx, articleResponse)
	if err != nil {
		log.Println("[Usecase][Article][InsertArticle] failed to insert article to elastic, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE, Error: err}
	}

	helper.FormatTimeArticleResponse(articleResponse)
	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: articleResponse, Error: nil}
}

func (u *usecase) UpdateArticle(ctx context.Context, req entity.UpdateArticleRequest) models.StandardResponseReq {
	var (
		article = entity.ArticleResponse{}
		err     error
	)

	article, err = u.repository.UpdateArticle(ctx, req)
	if err != nil {
		log.Println("[Usecase][Article][UpdateArticle] failed to update article, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_UPDATE_ARTICLE, Error: err}
	}

	// Update Article To Elasticsearch
	u.es.UpdateArticle(ctx, article)

	article.CreatedAt = helper.FormattedTime(article.CreatedAt)
	article.UpdatedAt = helper.FormattedTime(constant.TimeNow)

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_ARTICLE, Data: article, Error: nil}
}

func (u *usecase) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) models.StandardResponseReq {
	var (
		responseChan = make(chan models.StandardResponseReq)
		wg           sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := u.repository.DeleteArticle(ctx, req)
		if err != nil {
			log.Println("[Usecase][Article][DeleteArticle] failed to delete article, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_ARTICLE_POSTGRES, Error: err}
		}
	}()

	go func() {
		defer wg.Done()
		err := u.es.DeleteArticle(ctx, req)
		if err != nil {
			log.Println("[Usecase][Article][DeleteArticle] failed to delete article from elastic, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_ARTICLE_ELASTIC, Error: err}
		}
	}()

	go func() {
		wg.Wait()
		close(responseChan)
	}()

	for response := range responseChan {
		if response.Error != nil {
			return response
		}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_DELETE_ARTICLE, Error: nil}
}
