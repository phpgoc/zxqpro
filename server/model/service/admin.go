package service

import (
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/routes/request"
)

type AdminService struct {
	userDAO *dao.UserDAO
}

func NewAdminService(userDAO *dao.UserDAO) *AdminService {
	return &AdminService{
		userDAO: userDAO,
	}
}

func (s *AdminService) UpdatePassword(req request.AdminUpdatePassword) error {
	user, err := s.userDAO.GetByID(req.UserID)
	if err != nil {
		return err
	}
	user.Password = s.userDAO.Sha1Password(req.Password, user.ID)
	return s.userDAO.UpdatePassword(&user, user.Password)
}

func (s *AdminService) Create(req request.AdminRegister) error {
	user := entity.User{Name: req.Name, UserName: req.Name, Password: req.Password}
	err := s.userDAO.Create(&user)
	if err != nil {
		return err
	}
	password := s.userDAO.Sha1Password(user.Password, user.ID)

	return s.userDAO.UpdatePassword(&user, password)
}
