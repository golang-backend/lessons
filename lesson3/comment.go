package main

import "fmt"

/*
BU asosiy funksiya hisoblanadi va dastur shu yerda boshlanadi

	dasda
	dasd
	dasdad
*/
func main() {
	a := 0                      // a o'zgaruvchisi 0 dan katta qiymat qabul qiladi
	b := 0                      // b o'zgaruvchisi 0 dan kichik qiymat qabul qiladi
	fmt.Println("Natija:", a+b) // a va b ni qo'shib natijani chiqaradi
}

// A funksiyasi 2 ta sonni bir biriga qo'shadi va natijani qaytaradi
func A() {
	var num1 int // Birinchi son bu faqat musbat son bo'lishi kerak
}

// Ikki sonni bir biriga qo'shib beradi
// func add(a, b int) int {
// 	return a + b
// }

// 7-masaladagi int8 ning eng katta qiymatiga +1 qo'shib, go run qilib
// ko'ring.
// Natija nima bo'ldi? Bu hodisa overflow (to'lib ketish) deyiladi.
func task1() {
	
}