package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Product struct {
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Quantity int    `json:"qty,omitempty"`
	IsActive bool   `json:"is_active"`
	Brand    string `json:"-"`
}

type User struct {
	Name     string  `json:"name"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Address  Address `json:"address"`
}

type Address struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}

type Order struct {
	User      User      `json:"user"`
	Products  []Product `json:"products"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

func fullUpUser() User {
	user := User{
		Name:     "Umidjon",
		Login:    "umidjon",
		Password: "12345",
		Address: Address{
			Name:   "Samarqand",
			Target: "Registon",
		},
	}
	return user
}

func checkLogin(lgn string) bool {
	fmt.Printf("Login: ")
	var login string
	fmt.Scan(&login)

	if lgn != login {
		fmt.Printf("Login is invalid! Please try again\n")
		checkLogin(lgn)
	}

	return true
}

func checkPass(password string) bool {
	fmt.Printf("Password: ")
	var pass string
	fmt.Scan(&pass)

	if password != pass {
		fmt.Printf("Password is incorrect! Please try again\n")
		checkPass(password)
	}

	return true
}

func fullUpProducts() []Product {
	products := []Product{
		{
			Title:    "Iphone",
			Price:    1200,
			Quantity: 12,
			IsActive: true,
			Brand:    "Apple",
		},
		{
			Title:    "Samsung S23",
			Price:    1000,
			Quantity: 5,
			IsActive: true,
			Brand:    "Samsung",
		},
	}
	err := writeProducts(products)
	if err != nil {
		fmt.Println("fullUpProducts: ", err)
		return products
	}
	return products
}

func writeProducts(products []Product) error {
	products, err := getProducts()
	if err == nil && len(products) > 0 {
		return nil
	}

	file, err := fileCreate("products.json")
	// defer file.Close()
	if err != nil {
		fmt.Println("could not write products: ", err)
		return err
	}

	byteProducts, err := json.MarshalIndent(&products, "", "   ")
	if err != nil {
		fmt.Println("could not marshal products: ", err)
		return err
	}

	_, err = file.Write(byteProducts)
	if err != nil {
		fmt.Println("could not write product to file: ", err)
		return err
	}
	return nil
}

func getProducts() ([]Product, error) {
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error open file: ", err)
		return []Product{}, err
	}

	pByte, err := io.ReadAll(file)
	if err != nil {
		return []Product{}, err
	}
	products := []Product{}
	err = json.Unmarshal(pByte, &products)
	if err != nil {
		fmt.Println("could not unmarshal products: ", err)
		return []Product{}, err
	}

	return products, nil
}

func printProducts(products []Product) {
	count := 1
	for _, item := range products {
		fmt.Printf("%d. %s(%d) -> %d\n", count, item.Title, item.Quantity, item.Price)
		count++
	}
}

func main() {
	user := fullUpUser()

	fullUpProducts()

	checkLogin(user.Login)
	checkPass(user.Password)

	fmt.Println("---------------------MENU---------------------")
	fmt.Println("1. Product")
	fmt.Println("2. Profile")
	fmt.Println("3. Exit")

	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		products, err := getProducts()
		if err != nil {
			fmt.Println("something went wrong")
			return
		}
		printProducts(products)
	case 2:
		fmt.Println("Soon")
	case 3:
		fmt.Println("Bye Bye!!!")
		return
	}

	// product := Product{
	// 	Title:    "Computer",
	// 	Price:    1200,
	// 	Quantity: 12,
	// 	IsActive: true,
	// 	Brand:    "HP",
	// }

	// pByte, err := json.Marshal(&product)
	// if err != nil {
	// 	fmt.Println("could not marshal json: ", err.Error())
	// 	return
	// }

	// data := strings.Trim(string(pByte), "{")

	// if ok := json.Valid([]byte(data)); !ok {
	// 	fmt.Println("Invalid json")
	// 	return
	// }

	// file, err := os.Create("product.json")
	// if err != nil {
	// 	fmt.Println("could not create new file: ", err)
	// 	return
	// }

	// _, err = file.Write(pByte)
	// if err != nil {
	// 	fmt.Println("could not write product info: ", err.Error())
	// 	return
	// }

	// fmt.Println(string(pByte))

	// file.Close()
	// buf := new(bytes.Buffer)
	// enc := json.NewEncoder(buf)
	// dec := json.NewDecoder(buf)
	// enc.Encode(product)

	// var ptr Product
	// dec.Decode(&ptr)

	// fmt.Println(ptr)

}

func fileCreate(filename string) (*os.File, error) {
	var file *os.File
	if _, err := os.Stat(filename); err != nil {
		file, err = os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("could not create file: %v", err)
		}
	}
	return file, nil
}
