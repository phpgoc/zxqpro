package request

type AdminRegister struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
}
