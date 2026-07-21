package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	// understand string and byte

	// txt := "Hello World"
	// byt := []byte(txt)
	// fmt.Println("Txt: ", txt)
	// fmt.Println("Byte: ", byt)
	// fmt.Println("Length: ", len(txt))
	// // UTF-8

	// rus := "Хорошо это"
	// fmt.Println("Length RU: ", len(rus))
	// fmt.Println("Byte RU: ", []byte(rus))
	// fmt.Println("BYTE: ", []byte("Х"))

	// for _, v := range rus {
	// 	fmt.Printf("%v ", v)
	// }

	// Using strings
	// txt := "Stamotolik"

	// count := strings.Count(txt, "ta")
	// fmt.Println("t count: ", count)

	// exists := strings.Contains(txt, "g")
	// fmt.Println("g exists: ", exists)

	// begin := strings.HasPrefix(txt, "mo")
	// fmt.Println("mo begin: ", begin)

	// end := strings.HasSuffix(txt, "lik")

	// fmt.Println("lik end: ", end)

	// com := strings.Join([]string{"salom", "do'stim"}, "-")
	// fmt.Println(com)

	// repl := strings.Replace(txt, "o", "#", -1)
	// fmt.Println("Replaced: ", repl)

	// fmt.Println("Split: ", strings.Split("t$e$x$t", "$"))

	// fmt.Println("ToLower: ", strings.ToLower("Salom DunYO"))

	// Test byte

	// tx := []byte("Это байтовый срез")

	// fmt.Println(tx)

	// for i := range tx {
	// 	if tx[i]%2 == 0 {
	// 		tx[i]++
	// 	} else {
	// 		tx[i]--
	// 	}
	// }

	// fmt.Println(string(tx))

	// Using rune

	// ru := []rune("Это байтовый срез")

	// fmt.Println(ru)

	// txt := "dksdnka dmkcds"

	// for _, v := range txt {
	// 	if v == 'd' {
	// 		v = '*'
	// 	}
	// }

	// Unicode

	fmt.Println("IsDigit: ", unicode.IsDigit('0'))

	fmt.Println("IsLetter: ", unicode.IsLetter('u'))

	unicode.Is(unicode.Latin, 'R')

	unicode.ToLower('S') // s

	unicode.ToUpper('j') // J

	text := "读写汉字 - 学中文"
	fmt.Println("Length: ", len(text))
	fmt.Println(utf8.RuneCountInString(text))
}
