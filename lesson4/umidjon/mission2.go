package main
import "fmt"

// ikkinchi masala Taqqoslash
// bu funksiya istalgan uchta sondan kattasini aniqlaydi

func main() {
	var (
		a int 
		b int 
		c int
	)
	fmt.Scan(&a, &b, &c)
	fmt.Println("A, B, C", a, b, c)
	if a > b  && a > c  {   // agar a qolgan ikkita sondan katta bo‘lsa  eng katta son a
		fmt.Println("Big:", a)   
	} 
	if b > a  && b > c  {  // agar b qolgan ikkita sondan katta katta bo‘lsa eng katta b
		fmt.Println("Big:", b)
	} 
	if c > a  && c > b  {  // agar c qolgan ikkita sondan katta bo‘lsa eng katta son c
		fmt.Println("Big:", c)
	}
}