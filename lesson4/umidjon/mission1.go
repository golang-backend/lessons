package main
import "fmt"

// birinchi masala Toq || Juft
// bu funksiya sonni toq yoki juft ekanligini aniqlaydi
func main() {
    var a uint
	fmt.Scan(&a) // Scan bilam chop etilsa sonni o¯zimiz tanlaymiz
	fmt.Println("A:", a)
	if a % 2 == 0 {
		fmt.Println("Pair:", a)
	} else if a % 2 != 0 {  // nega else emas else if ishlatilmoqda chunki if dan keyin yana shart kiritilsa else if ishlatiladi kiritilmasa else 
		fmt.Println("Odd:", a)
	}
	// Tanlangan son agar ikkiga bo‘linsa juft bo♣linmasa toq  
}