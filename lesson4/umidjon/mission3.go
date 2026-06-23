
package main
import "fmt"

// uchinchi masala O‘quvchi bahosi
// bu funksiya oquvchining baliga qarab darajasini aniqlaydi
func main() {
	var ball int
	fmt.Scan(&ball)
	fmt.Println("Ball:", ball) 
	if ball >= 90 && ball <= 100 {
		fmt.Printf("Excelent: %d\n", ball) // Excelent - a‘lo
	}
	if ball >= 70 && ball <= 89 {
		fmt.Printf("Good: %d\n", ball) // good - yaxshi
	}
	if ball >= 50 && ball <= 69 {
		fmt.Printf("Satisfactory: %d\n", ball) // satisfactory - qoniqarli
	}
	if ball >= 0 && ball <= 49 {
		fmt.Printf("Unsatisfactory: %d\n", ball) // unsatisfactory - qoniqarsiz
	} 
    if ball > 100 || ball < 0 {
		fmt.Printf("Wrong ball: %d\n", ball)  // wrong ball - ball xato kiritildi
	}
}

