package main
import "fmt"
// beshinchi masala Iota yordamida statuslat va solishtirish 

// func main() {
// const (
// 	StatusYuborildi = iota  // 0
// 	StatusQabulqilindi      // 1
// 	StatusRadetildi         // 2
// )
//     joriyStatus := StatusYuborildi   // bunda joriyStatus := o degani
// 	x := joriyStatus != StatusRadetildi  // 0 && 2 teng ekanligi yolg‘onmi
// 	fmt.Println(x)
// }

// beshinchi masala  Yil fasli va harorat tavsiyasi
// bu funcsiya faslni aniqlab qanday kiyim kiyishni aytadi

// func main() {
// 	var month uint
// 	fmt.Scan(&month)
// 	fmt.Println("Month:", month)
// 	if month == 12 || month == 1 || month == 2 {
// 		fmt.Println("Winter: wear warm clothes") // Qish: issiq kiyim kiying
// 	}
// 	if month == 3 || month == 4 || month == 5 {  // Bahor: yengil kiyim kiying
// 		fmt.Println("Spring: wear light clothes")
// 	}
// 	if month == 6 || month == 7 || month == 8 {  // Yoz: futbolka kiying
// 		fmt.Println("Summer: wear T-shirts")
// 	}
// 	if month == 9 || month == 10 || month == 11 {
// 		fmt.Println("Autumn: wear light clothes")  // Kuz: yengil kiyim kiying
// 	}
// }


// beshinchi masala To‘g‘ri turtburchak 
// bu funksiya ikkita togri tortburchakdan yuzasi eng kattasi && perimetri teng || teng emasligini hisoblaydi

func main() {
	var (
		a uint 
		b uint 
		x uint
		y uint
	)
	fmt.Scan(&a, &b, &x, &y)
	fmt.Println("A:", "B:", "X:", "Y:", a, b, x, y)
	var s1 = a * b
	var s2 = x * y
	var p1 = 2 * a + 2 * b
	var p2 = 2 * x + 2 * y
    switch {
	case s1 > s2:
		fmt.Println("Eng katta yuza:", s1)
	case s1 < s2:
		fmt.Println("Eng katta yuza:", s2)	
	}
	switch {
	case p1 > p2:
		fmt.Println("Eng katta perimetr:", p1)
	case p1 < p2:
		fmt.Println("Eng katta perimetr:", p2)
	}
}
