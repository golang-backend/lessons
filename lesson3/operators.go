package main

import "fmt"

func main() {
	a := 24
	b := 23
	var c bool = a <= b
	// > < - qatiy
	// <= => - noqatiy
	// == tengligini tekshiradi
	// != teng emasligini tekshiradi

	var g bool = a == b
	var t = a != b

	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(t)

	var r = 90
	var v = a > b || a > r || r > b
	fmt.Println("==================")
	fmt.Println(v)

	h := 12
	var k = r > a && h > b

	fmt.Println("natija: ", k)

	var u = true
	fmt.Println("<--------------------->")
	fmt.Println(!u)
}
