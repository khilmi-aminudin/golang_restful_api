package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,max=200,min=1"`
}
