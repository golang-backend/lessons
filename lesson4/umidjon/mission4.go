package main

import "fmt"

// to‘rtinci masala Kalkulyator
// bu fuksiya sonlarni hisoblashda yordam beradi

func main() {
	// var a int
	// var b int
	// fmt.Scan(&a, &b)
	// fmt.Println("A:", a, "B:", b)

	// if b == 0 {
	// 	fmt.Println("Nolga bo‘lish mumkin emas:")
	// } else {
	// 	var result int = a / b
	// 	fmt.Println("Result:", result)
	// }

	var (
		a, b int
	)
	fmt.Printf("Birinchi sonni kirit: ")
	fmt.Scan(&a)
	fmt.Printf("Ikkinchi sonni kirit: ")
	fmt.Scan(&b)
	fmt.Println("Operatorni tanlang: ")
	fmt.Println("1 - '+' | 2 - '-' | 3 - '*' | 4 - '/'")
	var opt int
	fmt.Scan(&opt)
	if opt == 1 {
		fmt.Println("Natija: ", a+b)
	} else if opt == 2 {
		fmt.Println("Natija: ", a-b)
	} else if opt == 3 {
		fmt.Println("Natija: ", a*b)
	} else if opt == 4 && b != 0 {
		fmt.Println("Natija: ", a/b)
	} else {
		fmt.Println("Noto'g'ri qiymat kiritildi!!!")
	}
}
