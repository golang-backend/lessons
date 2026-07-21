package shape

import (
	"fmt"

	"lesson9/orders"
)

func BuildShape() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf(" *")
		}
		fmt.Print()
	}
}

func CorrectShape() {

}

func TestShape() {

}

func unitShape() {
	amount := orders.GetAmount()

	fmt.Println("Order Amount: ", amount)
}
