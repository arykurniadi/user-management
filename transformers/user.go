package transformers

import (
	"encoding/json"
	"fmt"

	"kriya.com/user-management/models"
)

type (
	User struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	UserList struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}

	UserCreate struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}

	UserUpdate struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}

	UserDelete struct {
		Id string `json:"id"`
	}
)

func (res *Transformer) TransformUser(item *models.User) {
	if item != nil {
		user := User{}

		var itemData map[string]interface{}
		err := json.Unmarshal([]byte(item.Data), &itemData)
		if err != nil {
		}

		var roleData map[string]interface{}
		err = json.Unmarshal([]byte(item.Role.Data), &roleData)
		if err != nil {
		}

		user.Id = item.Id.String()
		user.Username = fmt.Sprintf("%v", itemData["username"])
		user.Role = fmt.Sprintf("%v", roleData["roleName"])

		res.Data = user
	}
}

func (res *CollectionPagingTransformer) TransformUserList(arrUser []models.User, pagination *models.Pagination) {
	for _, item := range arrUser {
		var itemData map[string]interface{}
		err := json.Unmarshal([]byte(item.Data), &itemData)
		if err != nil {
			continue
		}

		user := UserList{}
		user.Id = item.Id.String()
		user.Username = fmt.Sprintf("%v", itemData["username"])
		user.Email = fmt.Sprintf("%v", itemData["email"])
		user.Status = fmt.Sprintf("%v", itemData["status"])

		res.Data = append(res.Data, user)
	}

	res.Meta = pagination
}

func (res *Transformer) TransformUserCreate(item *models.User) {
	if item != nil {
		var itemData map[string]interface{}
		err := json.Unmarshal([]byte(item.Data), &itemData)
		if err != nil {
		}

		user := UserCreate{}
		user.Id = item.Id.String()
		user.Username = fmt.Sprintf("%v", itemData["username"])
		user.Email = fmt.Sprintf("%v", itemData["email"])
		user.Status = fmt.Sprintf("%v", itemData["status"])

		res.Data = user
	}
}

func (res *Transformer) TransformUserUpdate(item *models.User) {
	if item != nil {
		var itemData map[string]interface{}
		err := json.Unmarshal([]byte(item.Data), &itemData)
		if err != nil {
		}

		user := UserUpdate{}
		user.Id = item.Id.String()
		user.Username = fmt.Sprintf("%v", itemData["username"])
		user.Email = fmt.Sprintf("%v", itemData["email"])
		user.Status = fmt.Sprintf("%v", itemData["status"])

		res.Data = user
	}
}

func (res *Transformer) TransformUserDelete(item *models.User) {
	if item != nil {
		user := UserDelete{}
		user.Id = item.Id.String()

		res.Data = user
	}
}
