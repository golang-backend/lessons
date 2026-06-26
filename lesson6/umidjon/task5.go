package main

import "fmt"

//
//

func main() {
	var x int
	fmt.Println("Nechta son kiritasan")
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Printf("Sonni kirting: ")
		var a int
		fmt.Scan(&a)
		if a <= 1 {
			fmt.Println("Tub son emas")
			continue
		}
		for j := 1; j < a; j++{
			if a % j != 0 {
				fmt.Println("Tub son")
				break
			}
		}
	}
}
