package main

import (
	"fmt"
	"strings"
)

// Masala 4. So'zlarni sanash va almashtirish
func main() {
	var txt = "My  name is    Umidjon"

	newText := strings.Split(txt, " ")
	new1 := strings.Fields(txt)

	fmt.Println("--> ", newText)
	fmt.Println("==> ", new1)
	fmt.Println("so'zlar soni: ", len(new1))
	newStr := strings.Join(new1, " ")
	fmt.Println("New text: ", newStr)
}
