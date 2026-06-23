package main

import "fmt"

func login() bool {
	var username, password string
	for i := 1; i <= 3; i++ {
		fmt.Print("Username: ")
		fmt.Scan(&username)

		fmt.Print("Password: ")
		fmt.Scan(&password)
		if username == "admin" && password == "12345" {
			fmt.Println("Xush kelibsiz")
			return true
		} else {
			fmt.Println("Xato!  qaytadan urinib koring")
		}
	}
	return false

}

type Mohsulot struct {
	ID    int
	Nomi  string
	Narxi float64
	Soni  int
}

type Sklad interface {
	Qoshish(m Mohsulot)
	Korish()
	NarxOzgartirish(id int, yangiNarx float64)
	Ochirish(id int)
}
type Ombor struct {
	mahsulotlar []Mohsulot
}

func (O *Ombor) Qoshish(m Mohsulot) {
	O.mahsulotlar = append(O.mahsulotlar, m)

}
func (o *Ombor) Korish() {
	for _,m := range o.mahsulotlar {
		fmt.Println(m)

	}
}
func (o *Ombor) NarxOzgartirish(id int, yangiNarx float64) {
	for i := range o.mahsulotlar {
		if o.mahsulotlar[i].ID == id {
			o.mahsulotlar[i].Narxi = yangiNarx
		}

	}
}
func (o *Ombor) Ochirish(id int) {
	for i := range o.mahsulotlar {
		if o.mahsulotlar[i].ID == id {
			o.mahsulotlar = append(o.mahsulotlar[:i], o.mahsulotlar[i+1:]...)
			break

		}

	}
}
func main() {
	if !login() {
		fmt.Println("3 marta urindingiz")
		return
	} else {
		fmt.Println("xush kelibsiz")
	}
	ombor := Ombor{
		mahsulotlar: []Mohsulot{
			{
				ID:    1,
				Nomi:  "Telefon",
				Narxi: 2500000,
				Soni:  10,
			},
			{
				ID:    2,
				Nomi:  "Noutbuk",
				Narxi: 8500000,
				Soni:  5,
			},
			{
				ID:    3,
				Nomi:  "Sichqoncha",
				Narxi: 150000,
				Soni:  20,
			},
		},
	}
	m := Mohsulot{
		ID:    5,
		Nomi:  "macbook",
		Narxi: 120000,
		Soni:  50,
	}
	ombor.Qoshish(m)
	ombor.Korish()
	ombor.NarxOzgartirish(1, 200000)
	fmt.Println(ombor)

}
