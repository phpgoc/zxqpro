package request

type AdminRegister struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
}

type AdminUpdatePassword struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required,min=8,complexPassword"`
}

type AdminCreateProject struct {
	Name        string `json:"name" binding:"required"`
	OwnerID     uint   `json:"owner_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}
