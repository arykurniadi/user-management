package user

import (
	"github.com/gin-gonic/gin"
	"kriya.com/user-management/models"
	"kriya.com/user-management/requests"
)

type IUserUsecase interface {
	GetUserById(c *gin.Context, id string) (*models.User, error)
	GetListUser(c *gin.Context, page int, perPage int) ([]models.User, *models.Pagination, error)
	Create(c *gin.Context, req requests.UserCreate) (*models.User, error)
	Update(c *gin.Context, req requests.UserUpdate) (*models.User, error)
	Delete(c *gin.Context, id string) (*models.User, error)
}
