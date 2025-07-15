package web

type CategoryCreateRequest struct {
	Name string `validate:"requeired,min=1,max=100" json:"name"`
}
