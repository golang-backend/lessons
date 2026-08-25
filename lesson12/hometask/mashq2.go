package main 
import "fmt"

// Masala 2. Satrni teskari o'girish

func main() {
	var x string
	fmt.Scan(&x)
	r := []rune(x)
	for i := len(r)-1; i >= 0; i-- {
		fmt.Printf("%c", r[i])
	}
}