package main

import "fmt"

func main() {
	// masala 7
	var x, p, y int

	fmt.Scan(&x, &p, &y)

	years := 0

	for x < y {
		x = x + x*p/100
		years++
	}

	fmt.Println(years)

	// masala 8
	var a, b string

	fmt.Scan(&a, &b)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				fmt.Println( a[i])
			}
		}
	}


}
