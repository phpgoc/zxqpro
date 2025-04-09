package routes

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/phpgoc/zxqpro/routes/middleware"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"

	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
)

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
func UserLogin(c *gin.Context) {
	var req request.UserLogin
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	user := entity.User{Name: req.Name}
	result := dao.Db.Where(user).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "用户不存在 或密码错误"))
		return
	}
	if user.Password != dao.Md5Password(req.Password, user.ID) {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "用户不存在 或密码错误"))
		return
	}

	cookie := generateCookie(user)
	for {
		_, have := interfaces.Cache.Get(cookie)
		if !have {
			break
		}
		cookie = generateCookie(user)
	}
	cookieStruct := pro_types.Cookie{ID: user.ID, UseMobile: req.UseMobile}

	if req.UseMobile {
		interfaces.Cache.Set(cookie, cookieStruct, cache.NoExpiration)
	} else {
		interfaces.Cache.Set(cookie, cookieStruct, 30*time.Minute)
	}

	c.SetCookie(utils.CookieName, cookie, 0, "/", "", false, true)
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// UserLogout  godoc
// @Summary user logout
// @Schemes
// @Description user logout
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/logout [post]
func UserLogout(c *gin.Context) {
	cookie, _ := c.Request.Cookie(utils.CookieName)
	interfaces.Cache.Delete(cookie.Value)
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// UserInfo  godoc
// @Summary user info
// @Schemes
// @Description user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonResponse{data=response.User} "成功响应"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	utils.LogInfo(fmt.Sprintf("%d", userId))
	user, err := dao.GetUserById(userId)
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

// UserUpdate  godoc
// @Summary user update
// @Schemes
// @Description user update1
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserUpdate true "UserUpdate"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/update [post]
func UserUpdate(c *gin.Context) {
	var req request.UserUpdate
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	user := entity.User{
		Email:    &req.Email,
		UserName: req.UserName,
		Avatar:   req.Avatar,
	}

	result := dao.Db.Model(entity.User{}).Where("id = ?", userId).Updates(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
}

// UserUpdatePassword  godoc
// @Summary user update_password
// @Schemes
// @Description user update_password
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserUpdatePassword true "UserUpdatePassword"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/update_password [post]
func UserUpdatePassword(c *gin.Context) {
	var req request.UserUpdatePassword
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	if req.OldPassword == req.NewPassword {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "新旧密码不能相同"))
		return
	}
	if req.NewPassword == req.NewPassword2 {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "两次新密码不一致"))
		return
	}
	user := entity.User{Password: dao.Md5Password(req.NewPassword, userId)}

	result := dao.Db.Model(entity.User{}).Where("id = ?", userId).Updates(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
}

// UserList  godoc
// @Summary user list
// @Schemes
// @Description user list
// @Tags User
// @Accept */*
// @Produce json
// @Param page query request.Page true "UserList"
// @Success 200 {object} response.CommonResponse[data=response.UserList] "成功响应"
// @Router /user/list [get]
func UserList(g *gin.Context) {
	var req request.Page
	if success := utils.ValidateQuery(g, &req); !success {
		return
	}
	var total int64
	dao.Db.Model(&entity.User{}).Count(&total)
	total = total - 1
	var responseUsers []response.User
	dao.Db.Model(entity.User{}).Where("deleted_at IS NULL").Where("id != ?", 1).Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Select("id, name, user_name, email, avatar").Find(&responseUsers)
	g.JSON(http.StatusOK, response.CreateResponse(0, "ok", response.UserList{
		Total: total,
		Users: responseUsers,
	}))
}

func generateCookie(user entity.User) string {
	// 生成 cookie 使用id， 当前时间戳，和一个随机8位字符串生成
	combined := fmt.Sprintf("%s%d%s", user.ID, time.Now().Unix(), utils.RandomString(8))
	hash := sha1.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}
