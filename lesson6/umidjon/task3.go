package main

import "fmt"

// uchinchi masala

func main() {
	var n int
	fmt.Println("Son kiriting")
	fmt.Scan(&n)
	fmt.Println("N:", n)
	var a = 0
	var b = 1
	for i := 0; i <= n; i++ {
		fmt.Println(a)
		c := a + b
		a = b
		b = c
	}
}
