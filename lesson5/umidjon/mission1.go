package main

import "fmt"

// birinvhi masala Solishtirish va narx
// // bu funkaiy ikkita a va b sonini taqqoslashda ishlatiladi
// func main() {
//     a := 15
// 	b := 10
// 	result := a >= b   //  natijada a soni b sonidan katta yoki tengligini rost yoki yogonligi tekshiriladi
// 	fmt.Println("Result", result)
// }

// birinchi masala Toq || Juft
// bu funksiya sonni toq yoki juft ekanligini aniqlaydi
// func main() {
//     var a uint
// 	fmt.Scan(&a) // Scan bilam chop etilsa sonni o¯zimiz tanlaymiz
// 	fmt.Println("A:", a)
// 	if a % 2 == 0 {
// 		fmt.Println("Pair:", a)
// 	} else if a % 2 != 0 {  // nega else emas else if ishlatilmoqda chunki if dan keyin yana shart kiritilsa else if ishlatiladi kiritilmasa else
// 		fmt.Println("Odd:", a)
// 	}
// 	// Tanlangan son agar ikkiga bo‘linsa juft bo♣linmasa toq
// }

// birinvhi masala Katta va kivhik sonni aniqlash
// bu fuksiya istalgan uvhta sonning ichidan eng kattasi va eng kichigini aniqlaydi
func main() {

	var	x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println("X:", "Y:", "Z:", x, y, z)
	
	if x > y && x > z {
		fmt.Println("Big number:", x)
	}
	if y > x && y > z {
		fmt.Println("Big number:", y)
	}
    if z > y && z > y {
		fmt.Println("Big number:", z)
	}
    if x < y && x < z {
		fmt.Println("Small number:", x)
	}
	if y < x && y < z {
		fmt.Println("Small number:", y)
	}
	if z < x && z < y {
		fmt.Println("Small number:", z)
	}
	if x == y || x ==z || z == y {
		fmt.Println("Aniqlanmadi:")
	}
}	
