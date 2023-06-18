package postgres

import (
	"context"
	"database/sql"
	"log"

	query "github.com/assyatier21/simple-cms-admin-v2/internal/repository/postgres/queries"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	msg "github.com/assyatier21/simple-cms-admin-v2/models/lib"
)

func (r *repository) GetCategoryTree(ctx context.Context) ([]entity.Category, error) {
	var (
		categories []entity.Category
		rows       *sql.Rows
		err        error
	)

	rows, err = r.db.Query(query.GET_CATEGORY_TREE)
	if err != nil {
		log.Println("[Repository][GetCategoryTree] failed to get list of categories, err: ", err)
		return nil, err
	}

	for rows.Next() {
		var temp = entity.Category{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetCategoryTree] failed to scan category, err :", err)
			return nil, err
		}
		categories = append(categories, temp)
	}

	if len(categories) == 0 {
		return []entity.Category{}, nil
	}

	return categories, nil
}
func (r *repository) GetCategoryDetails(ctx context.Context, id int) (entity.Category, error) {
	var (
		category entity.Category
		err      error
	)

	err = r.db.QueryRow(query.GET_CATEGORY_DETAILS, id).Scan(&category.Id, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetCategoryDetails] failed to scan category, err: ", err)
		return entity.Category{}, err
	}

	return category, nil
}
func (r *repository) InsertCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	err := r.db.QueryRow(query.INSERT_CATEGORY, category.Title, category.Slug, category.CreatedAt, category.UpdatedAt).Scan(&category.Id)
	if err != nil {
		log.Println("[Repository][InsertCategory] failed to insert category, err: ", err)
		return entity.Category{}, err
	}

	return category, nil
}
func (r *repository) UpdateCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	rows, err := r.db.Exec(query.UPDATE_CATEGORY, &category.Title, &category.Slug, &category.UpdatedAt, &category.Id)
	if err != nil {
		log.Println("[Repository][UpdateCategory] failed to update category, err: ", err)
		return entity.Category{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return entity.Category{}, nil
	}

	return category, nil
}
func (r *repository) DeleteCategory(ctx context.Context, id int) error {
	rows, err := r.db.Exec(query.DELETE_CATEGORY, id)
	if err != nil {
		log.Println("[Repository][DeleteCategory] failed to delete category, err: ", err)
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return msg.ERROR_NO_ROWS_AFFECTED
	}

	return nil
}
