package main
import (
	"fmt"
	c "unicode/utf8"
)

// Masala 1. So'z uzunligi va belgilar soni
func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println("baytlar soni: ", len(a))
	
	fmt.Println("belgilar soni: ", c.RuneCountInString(a))
}