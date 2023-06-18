package elasticsearch

import (
	"context"
	"encoding/json"
	"log"

	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	"github.com/olivere/elastic/v7"
)

func (r *elasticRepository) GetArticles(ctx context.Context, req entity.GetArticlesRequest) ([]entity.ArticleResponse, error) {
	var (
		articles = []entity.ArticleResponse{}
	)

	res, err := r.es.Search().Index(r.cfg.IndexArticle).From(req.Offset).Size(req.Limit).Sort(req.SortBy, req.OrderByBool).Do(ctx)
	if err != nil {
		return articles, err
	}

	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var article entity.ArticleResponse
			err = json.Unmarshal(hit.Source, &article)
			if err != nil {
				log.Println(err)
			}
			articles = append(articles, article)
		}
	}

	return articles, err
}

func (r *elasticRepository) GetArticleDetails(ctx context.Context, query elastic.Query) (entity.ArticleResponse, error) {
	var (
		article = entity.ArticleResponse{}
	)

	res, err := r.es.Search().Index(r.cfg.IndexArticle).Query(query).Do(ctx)
	if err != nil {
		return article, err
	}

	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &article)
		if err != nil {
			log.Println("Error failed to unmarshal json, err: ", err)
		}
	}

	return article, nil
}

func (r *elasticRepository) InsertArticle(ctx context.Context, article entity.ArticleResponse) error {
	var (
		articleJSON []byte
		err         error
	)

	articleJSON, err = json.Marshal(article)
	if err != nil {
		log.Println("[Elastic][InsertArticle] failed to marshal article, err: ", err)
		return err
	}

	_, err = r.es.Index().Index(r.cfg.IndexArticle).Id(article.ID).BodyJson(string(articleJSON)).Do(ctx)
	if err != nil {
		log.Println("[Elastic][InsertArticle] failed to insert article, err: ", err)
		return err
	}

	return err
}

func (r *elasticRepository) UpdateArticle(ctx context.Context, article entity.ArticleResponse) error {
	_, err := r.es.Update().Index(r.cfg.IndexArticle).Id(article.ID).Doc(article).Do(ctx)
	if err != nil {
		log.Println("[Elastic][UpdateArticle] failed to update article, err: ", err)
		return err
	}

	return nil
}

func (r *elasticRepository) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) error {
	_, err := r.es.Delete().Index(r.cfg.IndexArticle).Id(req.ID).Do(ctx)
	if err != nil {
		log.Println("[Elastic][DeleteArticle] failed to delete article, err: ", err)
		return err
	}

	return nil
}
