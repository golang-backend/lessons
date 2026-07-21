package main

import "fmt"

// RAM -> Random Access Memory -> Operativka
// HHD -> Hard Disk Drive
// SSD -> Solid State Drive

func main() {
	// a := 100
	// var b *int = &a

	// fmt.Println("a: ", a)
	// *b++
	// fmt.Println("a1: ", a)
	// fmt.Println("&a: ", &a)
	// fmt.Println("b: ", b)

	// ----------- new() --------------------

	// a := new(int)

	// b := a

	// fmt.Println("b: ", b)
	// fmt.Println("a: ", a)

	// one(a)

	// fmt.Println("a1:", a, *a)
	// fmt.Println("b1:", b, *b)

	a := 450

	b := &a
	// *b++
	c := &b
	fmt.Println("a: ", a, &a)
	fmt.Println("b: ", b, &b)
	fmt.Println("c:", c, &c)

	**c++
	fmt.Println("a1: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

// func one(x *int) {
// 	*x = 1
// }
