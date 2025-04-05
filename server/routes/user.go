package routes

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/phpgoc/zxqpro/middleware"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"

	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/request"
	"github.com/phpgoc/zxqpro/response"
)

// @BasePath /api

// UserRegister  godoc
// @Summary user register
// @Schemes
// @Description user register
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.Register true "UserRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/register [post]
func UserRegister(g *gin.Context) {
	var req request.Register
	if success := utils.Validate(g, &req); !success {
		return
	}
	user := entity.User{Name: req.Name, Password: req.Password, Email: req.Email}
	result := utils.Db.Create(&user)

	if result.RowsAffected == 1 {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
	password := md5Password(req.Password, user.ID)
	result = utils.Db.Model(&user).Update("password", password)
}

// UserLogin  godoc
// @Summary user login
// @Schemes
// @Description user login
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.Login true "UserRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var req request.Login
	if success := utils.Validate(c, &req); !success {
		return
	}
	user := entity.User{Name: req.Name}
	result := utils.Db.Where(user).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "用户不存在 或密码错误"))
		return
	}
	if user.Password != md5Password(req.Password, user.ID) {
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
	c.SetCookie(utils.CookieName, cookie, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// UserInfo  godoc
// @Summary user info
// @Schemes
// @Description user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonResponse{Data=response.User} "成功响应"
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
// @Param user body request.UpdateUser true "UserUpdate"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/update [post]
func UserUpdate(c *gin.Context) {
	var req request.UpdateUser
	if success := utils.Validate(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		UserName: req.UserName,
		Avatar:   req.Avatar,
	}

	result := utils.Db.Model(entity.User{}).Where("id = ?", userId).Updates(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
}

func md5Password(password string, id uint) string {
	combined := fmt.Sprintf("%s%d", password, id)
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(combined))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

func generateCookie(user entity.User) string {
	// 生成 cookie 使用id， 当前时间戳，和一个随机8位字符串生成
	combined := fmt.Sprintf("%s%d%s", user.ID, time.Now().Unix(), utils.RandomString(8))
	hash := sha1.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}
