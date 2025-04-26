package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/phpgoc/zxqpro/routes/response"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/pro_types"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/utils"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

//func (s *UserService) GetUserByID(id uint) (entity.User, error) {
//	user, err := s.userDAO.GetByID(id)
//	if err != nil {
//		return entity.User{}, err
//	}
//	return user, nil
//}

func (s *UserService) Login(c *gin.Context, name string, password string, longLogin bool) error {
	user, err := s.userDAO.GetByEntity(&entity.User{Name: name})
	if err != nil {
		return errors.New("用户不存在 或密码错误")
	}

	if user.Password != s.userDAO.Sha1Password(password, user.ID) {
		return errors.New("用户不存在 或密码错误")
	}

	cookie := generateCookie(user)
	for {
		have := interfaces.Cache.IsSet(cookie)
		if !have {
			break
		}
		cookie = generateCookie(user)
	}
	cookieStruct := pro_types.Cookie{ID: user.ID, LongLogin: longLogin}

	if longLogin {
		interfaces.Cache.Set(cookie, cookieStruct, cache.NoExpiration)
	} else {
		interfaces.Cache.Set(cookie, cookieStruct, 30*time.Minute)
	}

	c.SetCookie(my_runtime.CookieName, cookie, 0, "/", "", false, true)
	return nil
}

func (s *UserService) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie(my_runtime.CookieName)
	if err != nil {
		return
	}
	interfaces.Cache.Delete(cookie.Value)
	c.SetCookie(my_runtime.CookieName, "", -1, "/", "", false, true)
}

func (s *UserService) UpdatePassword(userID uint, req request.UserUpdatePassword) error {
	if req.OldPassword == req.NewPassword {
		return errors.New("新旧密码不能相同")
	}
	if req.NewPassword != req.NewPassword2 {
		return errors.New("两次新密码不一致")
	}
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		return err
	}
	if user.Password != s.userDAO.Sha1Password(req.OldPassword, userID) {
		return errors.New("旧密码错误")
	}
	return s.userDAO.UpdatePassword(&user, s.userDAO.Sha1Password(req.NewPassword, userID))
}

func (s *UserService) List(includeAdmin bool) (*response.UserList, error) {
	var res response.UserList
	var err error
	var users []entity.User

	users, err = s.userDAO.ListUsers(includeAdmin)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		res.List = append(res.List, response.User{
			ID:       user.ID,
			Name:     user.Name,
			UserName: user.UserName,
			Email:    user.Email,
			Avatar:   user.Avatar,
		})
	}

	return &res, nil
}

func generateCookie(user entity.User) string {
	// 生成 cookie 使用id， 当前时间戳，和一个随机8位字符串生成
	combined := fmt.Sprintf("%d%d%s", user.ID, time.Now().Unix(), utils.RandomString(8))
	hash := sha1.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}
