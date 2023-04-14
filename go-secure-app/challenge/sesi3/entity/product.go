package entity

type Product struct {
	ID          uint   `json:"id"`
	CreatedAt   string `json:"created_at, omitempty"`
	UpdatedAt   string `json:"updated_at, omitempty"`
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required."`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required."`
}