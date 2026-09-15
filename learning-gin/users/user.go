package users

import (
	"learning-gin/helper"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func CreateUser(req User) (User, error) {
	users, err := helper.ReadUsers()
	if err != nil {
		return User{}, err
	}

	req.ID = len(users) + 1
	req.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	users = append(users, req)

	if err := helper.SaveUsers(users); err != nil {
		return User{}, err
	}

	return req, nil
}

func GetAllUsers() ([]any, error) {
	users, err := helper.ReadUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (User, error) {
	users, err := helper.ReadUsers()
	if err != nil {
		return User{}, err
	}

	for _, u := range users {
		userMap, ok := u.(map[string]interface{})
		if !ok {
			continue
		}

		if int(userMap["id"].(float64)) == id {
			return User{
				ID:        int(userMap["id"].(float64)),
				Name:      userMap["name"].(string),
				Email:     userMap["email"].(string),
				CreatedAt: userMap["created_at"].(string),
			}, nil
		}
	}

	return User{}, nil
}
