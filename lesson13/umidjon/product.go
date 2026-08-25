package main

import "fmt"

func (u *User) getProducts(p []Mahsulotlar) {
	for _, item := range p {
		fmt.Printf("ID: %d %s Narxi: %.2f\n", item.Id, item.Nomi, item.Narxi)
	}
}

func (u *User) buyProduct(p []Mahsulotlar) error {
	fmt.Printf("mahsulotni tanlang: ")
	var productId int
	fmt.Scan(&productId)
	orders := []string{}
	for _, item := range p {
		if productId == item.Id {
			orders = append(orders, item.Nomi)
		}
	}

	u.SotibOlingan = orders

	return nil
}
