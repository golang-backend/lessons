package main

import (
	"errors"
	"fmt"
)

func (u *User) Auth(users []User) (*User, error) {
	var (
		login string
		pass  string
	)
	fmt.Printf("Login: ")
	fmt.Scan(&login)
	fmt.Printf("Parol: ")
	fmt.Scan(&pass)
	user := &User{}
	for _, item := range users {
		if item.Login == login && item.Parol == pass {
			u = &item
			user = u
			return user, nil
		}
	}

	return &User{}, errors.New("login yoki parol noto'g'ri!!!")
}

func (u *User) getProfile() {
	fmt.Printf("Ism: %s\nFamiliya: %s\n Balance:%.2f\n",
		u.Ism, u.Familiya, u.Balance,
	)
}
