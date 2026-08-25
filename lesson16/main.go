package main

import (
	"fmt"
	"log"
)

type Click struct {
	MerchantUserId string
	ApiKey         string
}

func (s Click) ClickPay(amount int) error {
	fmt.Printf("Click orqali %d to'landi.\n", amount)
	return nil
}

func (s Click) Pay(amount int) error {
	fmt.Printf("Click orqali %d to'landi.\n", amount)
	return nil
}

type Payme struct {
	MerchantId int
	SecretKey  string
}

func (s Payme) PaymePay(amount int) error {
	fmt.Printf("Payme orqali %d to'landi.\n", amount)
	return nil
}

func (s Payme) Pay(amount int) error {
	fmt.Printf("Payme orqali %d to'landi.\n", amount)
	return nil
}

type Payment interface {
	Pay(amount int) error
}

func NewPayment(payment Payment, amount int) {
	err := payment.Pay(amount)
	if err != nil {
		log.Fatalf("failed paymet: %v", err)
	}

	fmt.Println("Payment succeed")
}

func main() {
	// click := Click{
	// 	MerchantUserId: "12345",
	// 	ApiKey:         "JKjndekjn34eJKDNee3jc3erkn",
	// }

	// // // payment with Click
	// // err := click.ClickPay(1000)
	// // if err != nil {
	// // 	log.Fatalln("failed payment with click")
	// // }
	// NewPayment(click, 10000)

	// payme := Payme{
	// 	MerchantId: 122323,
	// 	SecretKey:  "JKjndekjn34eJKDNee3jc3erkn",
	// }

	// // // payment with Click
	// // err = payme.PaymePay(1000)
	// // if err != nil {
	// // 	log.Fatalln("failed payment with payme")
	// // }
	// NewPayment(payme, 20000)

	// empty interface
	// var a any
	// a = "salom"
	// a = 1233.23243

	// a = true
	// a = false
	// a = 12234
	universal("Salom")

	universal(1223334)

	universal(true)

	universal(3.14)

	user := User{
		name: "Umidjon",
	}

	universal(user)
}

type User struct {
	name string
}

func universal(gen interface{}) {
	if val, ok := gen.(int); ok {
		fmt.Println("Value is integer: ", val)
	}

	switch v := gen.(type) {
	case int:
		fmt.Println(v + 100)
	case string:
		fmt.Println("Value is string: ", v)
	case bool:
		fmt.Println("Value is boolean: ", v)
	case float64:
		fmt.Println("Value is float number: ", v)
	case User:
		fmt.Println("User: ", v)
	}

}


type umidjon uint

func (u umidjon) Salom() error {
	return  nil
}