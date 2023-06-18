package entity

type Category struct {
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Slug      string `json:"slug" form:"slug"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

type CategoryResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type GetCategoriesRequest struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

type GetCategoryDetailsRequest struct {
	ID int `json:"id" form:"id" validate:"required"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id" form:"id" validate:"required"`
}
