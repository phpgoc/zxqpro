package routes

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"

	"github.com/phpgoc/zxqpro/orm_model"
	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/request"
	"github.com/phpgoc/zxqpro/response"
)

// @BasePath /api

// UserRegister  godoc
// @Summary user register
// @Schemes
// @Description do hello
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
	user := orm_model.User{Name: req.Name, Password: req.Password, Email: req.Email}
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
	user := orm_model.User{Name: req.Name}
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

func md5Password(password string, id uint) string {
	combined := fmt.Sprintf("%s%d", password, id)
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(combined))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

func generateCookie(user orm_model.User) string {
	// 生成 cookie 使用id， 当前时间戳，和一个随机8位字符串生成
	combined := fmt.Sprintf("%s%d%s", user.ID, time.Now().Unix(), utils.RandomString(8))
	hash := sha1.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}
