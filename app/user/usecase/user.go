package usecase

import (
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	UserInterface "kriya.com/user-management/app/user"
	"kriya.com/user-management/models"
	"kriya.com/user-management/requests"
)

type UserUsecase struct {
	UserRepository UserInterface.IUserRepository
}

func NewUserUsecase(u UserInterface.IUserRepository) UserInterface.IUserUsecase {
	return &UserUsecase{
		UserRepository: u,
	}
}

func (a *UserUsecase) GetUserById(c *gin.Context, id string) (*models.User, error) {
	user, err := a.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *UserUsecase) GetListUser(c *gin.Context, page int, perPage int) ([]models.User, *models.Pagination, error) {
	users, pagination, err := a.UserRepository.GetListUser(page, perPage)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (a *UserUsecase) Create(c *gin.Context, req requests.UserCreate) (*models.User, error) {
	var (
		user models.User
	)

	userData, _ := json.Marshal(models.UserData{
		Username: req.Username,
		Email:    req.Email,
		Status:   "active",
	})

	user.Id = uuid.NewV4()
	user.Data = string(userData)

	roleId, err := a.UserRepository.CheckRoleName(req.Role)
	if err != nil {
		return nil, errors.New("Role name not found")
	}

	user.RoleId = roleId

	err = a.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserUsecase) Update(c *gin.Context, req requests.UserUpdate) (*models.User, error) {
	var (
		user models.User
	)

	userData, _ := json.Marshal(models.UserData{
		Username: req.Username,
		Email:    req.Email,
		Status:   "active",
	})

	id, _ := uuid.FromString(req.Id)

	user.Id = id
	user.Data = string(userData)

	roleId, err := a.UserRepository.CheckRoleName(req.Role)
	if err != nil {
		return nil, err
	}

	user.RoleId = roleId

	err = a.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserUsecase) Delete(c *gin.Context, id string) (*models.User, error) {
	row, err := a.UserRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return row, nil
}
