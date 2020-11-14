package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	userInterface "kriya.com/user-management/app/user"
	"kriya.com/user-management/models"
)

type UserRepository struct {
	ConnDB *gorm.DB
}

func NewUserRepository(ConnDB *gorm.DB) userInterface.IUserRepository {
	return &UserRepository{ConnDB}
}

func (m *UserRepository) GetUserById(id string) (*models.User, error) {
	tx := m.ConnDB

	var user models.User

	err := tx.Preload("Role").First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserRepository) GetListUser(page int, perPage int) ([]models.User, *models.Pagination, error) {
	tx := m.ConnDB

	var users []models.User

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	offset := (page * perPage) - perPage
	fmt.Println("page = ", page)
	fmt.Println("perPage = ", perPage)
	fmt.Println("offset = ", offset)

	_ = tx.Find(&users).Error
	total := len(users)

	err := tx.Preload("Role").Limit(perPage).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	pagination := models.BuildPagination(total, page, perPage, len(users))

	return users, pagination, nil
}

func (m *UserRepository) CheckRoleName(roleName string) (string, error) {
	tx := m.ConnDB

	var roles []models.Role

	err := tx.Find(&roles).Error
	if err != nil {
		return "", err
	}

	for _, item := range roles {
		var itemData map[string]interface{}
		err := json.Unmarshal([]byte(item.Data), &itemData)
		if err != nil {
			return "", err
			break
		}

		if itemData["roleName"] == roleName {
			return item.Id.String(), nil
			break
		}
	}

	return "", errors.New("Role name has not found")
}

func (m *UserRepository) Create(user models.User) error {
	tx := m.ConnDB

	err := tx.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *UserRepository) Update(item models.User) error {
	var user models.User

	tx := m.ConnDB

	err := tx.Where("id = ?", item.Id).First(&user).Error
	if err != nil {
		return err
	}

	user.Data = item.Data
	user.RoleId = item.RoleId
	tx.Save(&user)

	return nil
}

func (m *UserRepository) Delete(id string) (*models.User, error) {
	var user models.User

	tx := m.ConnDB

	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = tx.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
