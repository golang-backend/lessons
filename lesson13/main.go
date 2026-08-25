package main

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {
	// var a int
	// _, err := fmt.Scan(&a)
	// if err != nil {
	// 	fmt.Println("noto'g'ri qiymat kiritildi!! ", err)
	// 	return
	// } else {
	// 	fmt.Println("--> ", a)
	// }
	// res := divide(12, 0)
	// fmt.Println("DV Result: ", res)

	// r, err := IsText('3')
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// } else {
	// 	fmt.Println("-->> ", r)
	// }

	// xato := errors.New("user not found")

	// fmt.Println("ERR: ", xato)
	checkDefer()
}

func divide(a int, b int) int {
	if b == 0 {
		panic("0 ga bo'lib bo'lmaydi!!!")
	}
	return a / b
}

func IsText(str rune) (string, error) {
	if !unicode.IsLetter(str) {
		return "", errors.New("is not a string")
	}
	return string(str), nil
}

func checkDefer() {
	defer lastOne()
	defer fmt.Println("Program started")
	fmt.Println("Program is working")

	a := 12
	defer Test(a)
	a = 23
}

func lastOne() {
	fmt.Println("Oxirgi qiymat!!!")
}

func Test(a int) {
	fmt.Println("a: ", a)
}
