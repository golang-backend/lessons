package main

import (
	"fmt"
)

type User struct {
	Ism          string
	Familiya     string
	Login        string
	Parol        string
	Balance      float64
	SotibOlingan []string
}
type Mahsulotlar struct {
	Id    int
	Nomi  string
	Narxi float64
	Soni  int
}

func main() {
	k := []User{
		{
			Ism:          "Ali",
			Familiya:     "Aliyev",
			Login:        "ali",
			Parol:        "6975",
			Balance:      200000,
			SotibOlingan: []string{"Telefon", "Zaryadlash qurilmasi"},
		},
		{
			Ism:          "Vali",
			Familiya:     "Valiyev",
			Login:        "vali",
			Parol:        "1998",
			Balance:      300000,
			SotibOlingan: []string{"Telefon", "Sichqoncha"},
		},
		{
			Ism:          "Doli",
			Familiya:     "Valiyev",
			Login:        "doli",
			Parol:        "2000",
			Balance:      700000,
			SotibOlingan: []string{"Telivizor", "Sichqoncha"},
		},
	}
	m := []Mahsulotlar{
		{
			Id:    1,
			Nomi:  "Telivizor",
			Narxi: 250000,
			Soni:  10,
		},
		{
			Id:    2,
			Nomi:  "Sichqoncha",
			Narxi: 20000,
			Soni:  25,
		},
		{
			Id:    3,
			Nomi:  "Tlefon",
			Narxi: 200000,
			Soni:  15,
		},
	}

	// call to auth function

	user := &User{}
	for {
		profile, authErr := user.Auth(k)
		if authErr != nil {
			fmt.Println(authErr.Error())
			continue
		} else {
			user = profile
			break
		}
	}
	fmt.Println("--->> ", user)
	// call to home menu
	user.Home(m)
}

func (u *User) Home(p []Mahsulotlar) {
	fmt.Println("1. Mahsulotlarni ko'rish.")
	fmt.Println("2. Mahsulot sotib olish.")
	fmt.Println("3. Profile.")
	fmt.Println("4. Chiqish.")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		u.getProducts(p)
	case 2:
		u.buyProduct(p)
	case 3:
		u.getProfile()
	}
}
