package main

import "fmt"

type Person struct { // structure
	name      string
	age       int
	email     string
	isStudent bool
}

// func main() {
// var (
// 	name  string
// 	age   int
// 	email string
// )
// name = "John"
// age = 20
// email = "john@gmail.com"
// fmt.Println(name, age, email)

// declare Struct
// p := Person{
// 	name:      "Tomy",
// 	age:       23,
// 	email:     "tomy@gmail.com",
// 	isStudent: false,
// }

// fmt.Println(p)
// fmt.Println("Name: ", p.name)
// fmt.Printf("Age: %d\n", p.age)

// 	k := Hisob{
// 		Owner:   "Umidjon",
// 		Balance: 100,
// 	}

// 	fmt.Println(k.getBalance())

// 	// k.addBalance(50)
// 	// k.Balance = k.Balance + 50
// 	k.Balance = 150

// 	fmt.Println(k.getBalance())

// 	testHisob(&k)
// }

type Hisob struct {
	Owner   string
	Balance int64
}

func (h Hisob) getBalance() int64 {
	return h.Balance
}

func (h *Hisob) addBalance(summa int64) {
	h.Balance = h.Balance + summa
}

func testHisob(h *Hisob) int64 {
	return h.Balance
}

type Human struct {
	Id    int
	Phone string
	Address
}

type Address struct {
	Name   string
	Target string
}

func (a Address) getAddress() string {
	return a.Name + ", " + a.Target
}

func main() {
	odam := Human{
		Id:    12,
		Phone: "98232323434",
		Address: Address{
			Name:   "Tashkent, Yakkasaroy 13A",
			Target: "NestOne hotel",
		},
	}

	fmt.Println(odam.getAddress())

}
