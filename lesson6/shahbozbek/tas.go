package main

import "fmt"

func main() {
	// 1- masala
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i

	}
	fmt.Println("Yigindi->", sum)

	// 2- masala
	kopaytma := 1
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			kopaytma = kopaytma * i
		}
	}
	fmt.Println("1 dan 50 gacha juft sonlar kopytmasi ->", kopaytma)

	// 3 -masala
	// fibanacchi ketma ketligi
	n := 8
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b

	}
	
}
