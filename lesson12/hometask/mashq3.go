package main

import (
	"bufio"
	"fmt"
	"os"
	y "unicode"
)

// Masala 3. Unli va undosh harflarni sanash

func main() {
	txt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	caunt := 0
	sum := 0
	r := []rune(txt)
	for i := 0; i < len(r); i++ {
		if !y.IsLetter(r[i]) {
			continue
		}
		b := string(y.ToLower(r[i]))
		if b == "a" || b == "i" || b == "o" || b == "u" || b == "e" {
			caunt++
		} else {
			sum++
		}

	}
	fmt.Printf("Unli harflar %d ta\n", caunt)
	fmt.Printf("Undosh harflar %d ta\n", sum)
}
