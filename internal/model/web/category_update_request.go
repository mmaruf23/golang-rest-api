package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=20,min=1" json:"name"`
}
