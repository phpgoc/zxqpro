package request

type UserLogin struct {
	Name      string `json:"name" binding:"required" default:"admin"`
	Password  string `json:"password" binding:"required" default:"Aa123456"`
	LongLogin bool   `json:"long_login" binding:"omitempty"`
}

type UserUpdate struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   byte   `json:"avatar" binding:"required,min=0,max=20"`
}

type UserUpdatePassword struct {
	OldPassword  string `json:"old_password" binding:"required"`
	NewPassword  string `json:"new_password" binding:"required,min=8,complexPassword"`
	NewPassword2 string `json:"new_password2" binding:"required,min=8,complexPassword"`
}

type UserList struct {
	ProjectID    uint `form:"project_id" binding:"min=0"`
	IncludeAdmin bool `form:"include_admin"`
}
