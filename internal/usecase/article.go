package usecase

// func (u *usecase) GetArticles(ctx context.Context, req entity.GetArticlesRequest) models.StandardResponseReq {
// 	var (
// 		articles = []entity.ArticleResponse{}
// 	)

// 	req.SortBy = helper.ValidateSortBy(req.SortBy)
// 	req.OrderByBool = helper.ValidateOrderBy(req.OrderBy)

// 	articles, err := u.es.GetArticles(ctx, req)
// 	if err != nil {
// 		log.Println("[Usecase][GetCategoryTree] failed to get list of articles, err: ", err)
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLES, Error: err}
// 	}

// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: articles, Error: nil}
// }

// func (u *usecase) GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) models.StandardResponseReq {
// 	var (
// 		article = entity.ArticleResponse{}
// 		query   elastic.Query
// 	)

// 	query = elastic.NewMatchQuery(constant.ID, req.ID)
// 	article, err := u.es.GetArticleDetails(ctx, query)
// 	if err != nil {
// 		log.Println("[Usecase][GetArticleDetails] failed to get article details, err: ", err)
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLE_DETAILS, Error: err}
// 	}

// 	helper.FormatTimeArticleResponse(&article)
// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: article, Error: nil}
// }

// func (u *usecase) InsertArticle(ctx context.Context, req entity.InsertArticleRequest) models.StandardResponseReq {
// 	reqArticle := entity.Article{
// 		ID:          helper.GenerateUUIDString(),
// 		Title:       req.Title,
// 		Slug:        req.Slug,
// 		HTMLContent: req.HTMLContent,
// 		CategoryID:  req.CategoryID,
// 		CreatedAt:   constant.TimeNow,
// 		UpdatedAt:   constant.TimeNow,
// 	}

// 	err := json.Unmarshal([]byte(req.Metadata), &reqArticle.MetaData)
// 	if err != nil {
// 		log.Println("[Usecase][InsertArticle] failed to unmarshal article metadata, err: ", err)
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE, Error: err}
// 	}

// 	article, err := u.repository.InsertArticle(ctx, reqArticle)
// 	if err != nil {
// 		log.Println("[Usecase][InsertArticle] failed to insert article, err: ", err)
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE, Error: err}
// 	}

// 	err = u.es.InsertArticle(ctx, article)
// 	if err != nil {
// 		log.Println("[Usecase][InsertArticle] failed to insert article to elastic, err: ", err)
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE, Error: err}
// 	}

// 	helper.FormatTimeArticleResponse(&article)
// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: article, Error: nil}
// }

// func (u *usecase) UpdateArticle(ctx context.Context, req entity.UpdateArticleRequest) models.StandardResponseReq {
// 	var (
// 		article entity.ArticleResponse
// 		err     error
// 	)

// 	article, err = u.repository.GetArticleDetails(ctx, entity.GetArticleDetailsRequest{ID: req.ID})
// 	if err != nil {
// 		log.Println("[Usecase][UpdateArticle] failed to get article details, err: ", err)

// 	}

// 	if req.Title != "" {
// 		article.Title = req.Title
// 	}

// 	if req.Slug != "" {
// 		article.Slug = req.Slug
// 	}

// 	if req.HTMLContent != "" {
// 		article.HTMLContent = req.HTMLContent
// 	}

// 	if req.CategoryID != 0 {
// 		article.ResCategory.ID = req.CategoryID
// 	}

// 	if req.Metadata != "" {
// 		err = json.Unmarshal([]byte(req.Metadata), &article.MetaData)
// 		if err != nil {
// 			log.Println("[Usecase][UpdateArticle] failed to update article, err: ", err)

// 		}
// 	}

// 	article, err = u.repository.UpdateArticle(ctx, entity.Article{
// 		ID:          req.ID,
// 		Title:       article.Title,
// 		Slug:        article.Slug,
// 		HTMLContent: article.HTMLContent,
// 		CategoryID:  article.ResCategory.ID,
// 		MetaData:    article.MetaData,
// 		CreatedAt:   article.CreatedAt,
// 		UpdatedAt:   helper.FormattedTime(constant.TimeNow),
// 	})

// 	if err != nil {
// 		log.Println("[Usecase][UpdateArticle] failed to update article, err: ", err)

// 	}

// 	u.es.UpdateArticle(ctx, article)

// 	article.CreatedAt = helper.FormattedTime(article.CreatedAt)
// 	article.UpdatedAt = helper.FormattedTime(constant.TimeNow)

// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_ARTICLE, Data: article, Error: nil}
// }

// func (u *usecase) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) error {
// 	var (
// 		articleDeleted bool
// 		elasticDeleted bool
// 		err            error
// 		wg             sync.WaitGroup
// 	)

// 	wg = sync.WaitGroup{}
// 	wg.Add(2)

// 	go func() {
// 		err := u.repository.DeleteArticle(ctx, req)
// 		if err != nil {
// 			log.Println("[Usecase][DeleteArticle] failed to delete article, err: ", err)
// 		} else {
// 			articleDeleted = true
// 		}
// 		wg.Done()
// 	}()

// 	go func() {
// 		err := u.es.DeleteArticle(ctx, req)
// 		if err != nil {
// 			log.Println("[Usecase][DeleteArticle] failed to delete article from elastic, err: ", err)
// 		} else {
// 			elasticDeleted = true
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	if !articleDeleted && !elasticDeleted {
// 		err = errors.New("article not found")
// 		return err
// 	}

// 	return nil
// }
