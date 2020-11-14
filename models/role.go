package models

import uuid "github.com/satori/go.uuid"

type (
	Role struct {
		Id   uuid.UUID `json:"id" gorm:"primary_key,column:id"`
		Data string    `json:"data" gorm:"column:data"`
	}

	RoleData struct {
		RoleName string `json:"roleName"`
	}
)

func (Role) TableName() string {
	return "roles"
}
