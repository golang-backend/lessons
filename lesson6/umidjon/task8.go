package main

import "fmt"

// sakkizinvhi masala Musbat sonlar o‘rtachasi

func main() {
	var n int
	fmt.Println("Nechta son kiritasiz")
	fmt.Scan(&n)
	var sum = 0
	for i := 0; i < n; i++ {
		fmt.Println("Sonni kiriting")
		var x int
		fmt.Scan(&x)
		if x < 0 {
			fmt.Println("Musnat son kiriting")
			continue
		}
		sum += x
	}
	var a = sum / n
	fmt.Println("O‘rta qiymat:", a)
}
