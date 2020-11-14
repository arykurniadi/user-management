package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"kriya.com/user-management/config"
	gorm "kriya.com/user-management/db"

	routes "kriya.com/user-management/app"

	HCRepository "kriya.com/user-management/app/health-check/repository"

	UserRepository "kriya.com/user-management/app/user/repository"

	HCUsecase "kriya.com/user-management/app/health-check/usecase"
	UserUsecase "kriya.com/user-management/app/user/usecase"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()

	r.Use(gin.Recovery())

	hcr := HCRepository.NewHealthCheckRepository(db)
	userRepo := UserRepository.NewUserRepository(db)

	hcu := HCUsecase.NewHealthCheckUsecase(hcr)
	user := UserUsecase.NewUserUsecase(userRepo)

	routes.HealthCheckHttpHandler(r, hcu)
	routes.UserHttpHandler(r, user)

	// check auth endpoint for admin access
	r.Use(AuthMiddleware())
	routes.AdminHttpHandler(r, user)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth == true && username != "admin" && password != "admin" {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}
