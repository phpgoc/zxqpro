package request

type Register struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
	Email    string `json:"email" binding:"omitempty,email"`
}
