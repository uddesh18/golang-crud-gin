package request

type CreateTagRequest struct {
	Name string `json:"name" validate:"required, min=1 , max=10"`
}
