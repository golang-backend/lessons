package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Masala 5. Palindrom tekshiruvchi
// Satr palindrom ekanligini tekshiramiz. Bo'sh joy, tinish belgilari va
// katta-kichik harf farqini e'tiborga olmaymiz. Kirill harflari uchun ham ishlaydi.

func palindrom(satr string) bool {
	var toza []rune
	for _, r := range satr {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			toza = append(toza, unicode.ToLower(r))
		}
	}

	// Checking with two side
	left := 0
	right := len(toza) - 1
	for left < right {
		if toza[left] != toza[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Satrni kiriting: ")
	satr, _ := reader.ReadString('\n')
	satr = strings.TrimSpace(satr)

	if palindrom(satr) {
		fmt.Println("Ha, bu palindrom ✅")
	} else {
		fmt.Println("Yo'q, bu palindrom emas ❌")
	}
}
