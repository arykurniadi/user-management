package app

import (
	"github.com/gin-gonic/gin"

	HealthCheckInterface "kriya.com/user-management/app/health-check"
	UserInterface "kriya.com/user-management/app/user"

	HCHandler "kriya.com/user-management/app/health-check/handler"
	UserHandler "kriya.com/user-management/app/user/handler"
)

func HealthCheckHttpHandler(r *gin.Engine, us HealthCheckInterface.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}

	route := r.Group("/test")
	route.GET("/health-check", handler.Check)
}

func UserHttpHandler(r *gin.Engine, us UserInterface.IUserUsecase) {
	handler := &UserHandler.UserHandler{
		UserUsecase: us,
	}

	route := r.Group("/user")
	route.GET("/", handler.GetUserById)
	route.GET("/list", handler.GetListUser)
}

func AdminHttpHandler(r *gin.Engine, us UserInterface.IUserUsecase) {
	handler := &UserHandler.UserHandler{
		UserUsecase: us,
	}

	route := r.Group("/admin")
	route.POST("/create", handler.Create)
	route.POST("/update", handler.Update)
	route.GET("/delete", handler.Delete)
}
