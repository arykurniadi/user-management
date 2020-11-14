package user

import (
	"kriya.com/user-management/models"
)

type IUserRepository interface {
	GetUserById(string) (*models.User, error)
	GetListUser(int, int) ([]models.User, *models.Pagination, error)
	CheckRoleName(roleName string) (string, error)
	Create(models.User) error
	Update(models.User) error
	Delete(string) (*models.User, error)
}
