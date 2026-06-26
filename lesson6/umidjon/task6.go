package main
import "fmt"
// oltinchi masala Kopaytirish jadvali


func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			x := i*j
			fmt.Printf("|%4d|", x)
		}
		fmt.Println()
	}
}