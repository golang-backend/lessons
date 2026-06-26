package main

import "fmt"

// to‘rtinchi masala Raqamlar yig‘indisi
//

func main() {
	var a = 123456
	fmt.Println("A:", a)
	var sum = 0
	for a > 0 {
		x := a % 10
		sum += x
		a = a / 10
	}
	fmt.Println("Raqqamlar yigindisi:", sum)
}
