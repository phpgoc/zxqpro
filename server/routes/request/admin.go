package request

type AdminRegister struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
}

type AdminUpdatePassword struct {
	ID       uint   `json:"id" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
}
