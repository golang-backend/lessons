package main
import "fmt"
// ikkinchi masala Dokon balansi && narx
// bu funksiya balans maxsulotni sotib olishga yetarli || yetarli emasligini aniqlaydi 
// func main() {
// 	balance := 45000  //  45000 mijozning balansi
// 	charge := 50000  // 50000 mahsulot narxi
// 	result := balance >= charge 
// 	fmt.Println("Result",result)
// }

// ikkinchi masala Taqqoslash
// bu funksiya istalgan uchta sondan kattasini aniqlaydi

// func main() {
// 	var (
// 		a int 
// 		b int 
// 		c int
// 	)
// 	fmt.Scan(&a, &b, &c)
// 	fmt.Println("A, B, C", a, b, c)
// 	if a > b  && a > c  {   // agar a qolgan ikkita sondan katta bo‘lsa  eng katta son a
// 		fmt.Println("Big:", a)   
// 	} 
// 	if b > a  && b > c  {  // agar b qolgan ikkita sondan katta katta bo‘lsa eng katta b
// 		fmt.Println("Big:", b
// 	} 
// 	if c > a  && c > b  {  // agar c qolgan ikkita sondan katta bo‘lsa eng katta son c
// 		fmt.Println("Big:", c)
// 	}
// }



// ikkinchi masala Juft || toq
// bu funksiya istalgan 5 xonali sonni juft oki toq ligini aniqlaydi

func main() {
	var a uint
	var b uint
	var c uint
	var d uint
	var f uint
	fmt.Scan(&a, &b, &c, &d, &f)
	fmt.Println("A:", "B:", "D:", "F:", a, b, c, d, f)
	var number = 10000 * a + 1000 * b + 100 * c + 10 * d + f
    fmt.Println("Number:", number)
	switch {
	case f % 2 == 0:
		fmt.Println("Pair number:", number)
	case f % 2 != 0:
		fmt.Println("Odd number:", number)    
	}
	if a % 2 == 0 && a != 0 {
		fmt.Println("Pair number:", a)
	}
	if a % 2 != 0 && a != 0 {
		fmt.Println("Odd number:", a)
	}
	if b % 2 == 0 {
		fmt.Println("Pair number:", b)
	}
	if b % 2 != 0 {
		fmt.Println("Odd number:", b)
	}
    if c % 2 == 0 {
		fmt.Println("Pair number:", c)
	}
	if c % 2 != 0 {
		fmt.Println("Odd number:", c)
	}
	if d % 2 == 0 {
		fmt.Println("Pair number:", d)
	}
	if d % 2 != 0 {
		fmt.Print("Odd number:", d)
	}
	if f % 2 == 0 {
		fmt.Println("Pair number:", f)
	}
	if f % 2 != 0 {
		fmt.Println("Odd number:", f)
	}
}