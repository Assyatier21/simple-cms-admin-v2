package postgres

import (
	"context"
	"log"

	"github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres/queries"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
)

func (r *repository) GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) ([]entity.Category, error) {
	var (
		categories []entity.Category
	)

	rows, err := r.db.Query(queries.GET_CATEGORY_TREE, req.Limit, req.Offset)
	if err != nil {
		log.Println("[Repository][GetCategories] failed to query categories, err: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Println("[Repository][GetCategories] failed to scan category, err: ", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repository) GetCategoryDetails(ctx context.Context, req entity.GetCategoryDetailsRequest) (entity.Category, error) {
	var (
		category entity.Category
		err      error
	)

	err = r.db.QueryRow(queries.GET_CATEGORY_DETAILS, req.ID).Scan(&category.ID, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetCategoryDetails] failed to scan category, err: ", err)
		return entity.Category{}, err
	}

	return category, nil
}

func (r *repository) InsertCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	err := r.db.QueryRow(queries.INSERT_CATEGORY, category.Title, category.Slug, category.CreatedAt, category.UpdatedAt).Scan(&category.ID)
	if err != nil {
		log.Println("[Repository][InsertCategory] failed to insert category, err: ", err)
		return entity.Category{}, err
	}

	return category, nil
}

func (r *repository) UpdateCategory(ctx context.Context, category entity.Category) error {
	_, err := r.db.Exec(queries.UPDATE_CATEGORY, &category.Title, &category.Slug, &category.UpdatedAt, &category.ID)
	if err != nil {
		log.Println("[Repository][UpdateCategory] failed to update category, err: ", err)
		return err
	}

	return nil
}

func (r *repository) DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) error {
	_, err := r.db.Exec(queries.DELETE_CATEGORY, req.ID)
	if err != nil {
		log.Println("[Repository][DeleteCategory] failed to delete category, err: ", err)
		return err
	}

	return nil
}
