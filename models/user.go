package models

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		Id     uuid.UUID `json:"id" gorm:"primary_key,column:id"`
		Data   string    `json:"data" gorm:"column:data"`
		RoleId string    `gorm:"column:role_id"`
		Role   Role      `gorm:"foreignKey:RoleId;association_foreignkey:id"`
	}

	UserData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}
)

func (User) TableName() string {
	return "users"
}
