package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Masala 4. So'zlarni sanash va almashtirish
// Foydalanuvchidan gap qabul qilamiz, so'zlar sonini aniqlaymiz va
// ortiqcha bo'sh joylarni bitta bo'sh joyga aylantiramiz.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Gapni kiriting: ")
	text, _ := reader.ReadString('\n')

	// Satr oxiridagi yangi qatorni olib tashlaymiz
	text = strings.TrimSpace(text)

	// strings.Fields textni bo'sh joylar bo'yicha bo'lib beradi va
	// ketma-ket kelgan bo'sh joylarni o'zi hisobga olmaydi.
	words := strings.Fields(text)

	// So'zlarni bitta bo'sh joy bilan qayta birlashtiramiz -> tozalangan gap.
	tozalangan := strings.Join(words, " ")

	fmt.Println("Tozalangan gap:", tozalangan)
	fmt.Println("So'zlar soni:", len(words))
}
