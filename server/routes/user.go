package routes

import (
	"net/http"

	"github.com/phpgoc/zxqpro/model/service"

	"github.com/phpgoc/zxqpro/my_runtime"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"

	"github.com/phpgoc/zxqpro/interfaces"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @BasePath /api

// UserLogin  godoc
// @Summary user login
// @Schemes
// @Description user login
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserLogin true "UserRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/login [post]
func (h *UserHandler) UserLogin(c *gin.Context) {
	var req request.UserLogin
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	err := h.userService.Login(c, req.Name, req.Password, req.LongLogin)
	if err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// Logout  godoc
// @Summary user logout
// @Schemes
// @Description user logout
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/logout [post]
func (h *UserHandler) Logout(c *gin.Context) {
	cookie, _ := c.Request.Cookie(my_runtime.CookieName)
	userId := service.GetUserIDFromAuthMiddleware(c)
	interfaces.Cache.Delete(cookie.Value)
	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	sseManager.UnregisterUser(userId)
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// Info  godoc
// @Summary user info
// @Schemes
// @Description user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonResponse{data=response.User} "成功响应"
// @Router /user/info [get]
func (h *UserHandler) Info(c *gin.Context) {
	userID := service.GetUserIDFromAuthMiddleware(c)
	user, err := dao.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	responseStruct := response.User{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", responseStruct))
}

// Update  godoc
// @Summary user update
// @Schemes
// @Description user update1
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserUpdate true "Update"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/update [post]
func (h *UserHandler) Update(c *gin.Context) {
	var req request.UserUpdate
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	user := entity.User{
		Email:    &req.Email,
		UserName: req.UserName,
		Avatar:   req.Avatar,
	}

	result := my_runtime.Db.Model(entity.User{}).Where("id = ?", userID).Updates(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
}

// UpdatePassword  godoc
// @Summary user update_password
// @Schemes
// @Description user update_password
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserUpdatePassword true "UpdatePassword"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/update_password [post]
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req request.UserUpdatePassword
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	if err := h.userService.UpdatePassword(userID, req); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// List  godoc
// @Summary user list
// @Schemes
// @Description user list
// @Tags User
// @Accept */*
// @Produce json
// @Param page query request.UserList true "List"
// @Success 200 {object} response.CommonResponse[data=response.UserList] "成功响应"
// @Router /user/list [get]
func (h *UserHandler) List(c *gin.Context) {
	var req request.UserList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	list, err := h.userService.List(req.IncludeAdmin)
	if err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", list))
}
