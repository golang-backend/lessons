package main

import (
	"encoding/json"
	"fmt"
)

type Kitob struct {
	Name     string
	Author   string
	Quantity int
}

type User struct {
	Name  string
	Books []Kitob
}

type Library struct {
	Users []*User
	Books []*Kitob
}

func main() {
	lib := Library{
		Users: []*User{},
		Books: []*Kitob{
			{
				Name:     "New Book1",
				Author:   "Author1",
				Quantity: 10,
			},
			{
				Name:     "Book2",
				Author:   "Author2",
				Quantity: 5,
			},
		},
	}

	err := lib.KitobOlish("User", "Book2")
	if err != nil {
		fmt.Println("Kitop topilmadi!!!")
		return
	}
	t, _ := json.MarshalIndent(&lib, "", "   ")
	fmt.Println(string(t))
}

func (r *Library) KitobOlish(userName, bookName string) error {
	count := 0
	for _, item := range r.Books {
		if item.Name == bookName {
			item.Quantity--
			r.Users = append(r.Users, &User{
				userName,
				[]Kitob{*item},
			})
			return nil
		} else {
			count++
		}
	}

	if count > 0 {
		return fmt.Errorf("Kitop topilmadi")
	}
	return nil
}
