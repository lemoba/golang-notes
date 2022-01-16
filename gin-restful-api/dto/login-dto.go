package dto

// LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:3"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
