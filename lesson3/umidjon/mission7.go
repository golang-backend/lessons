package main

import "fmt"

// yettinchi masala Kabisa yili sharti (mantiqiy savol)

// bu funksiya yil kabisa yoki kabisa emasligini aniqlaydi
func main() {
	const year = 2024
	x := year%4 == 0 && year%100 != 0 || year%400 == 0
	fmt.Println("Kabisa yiligi rostmi:", x)

}
