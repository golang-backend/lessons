package main

import "fmt"

// yettinchi masala Kabisa yili sharti (mantiqiy savol)

// bu funksiya yil kabisa yoki kabisa emasligini aniqlaydi
// func main() {
// 	const year = 2024
// 	x := year%4 == 0 && year%100 != 0
// 	fmt.Println("Kabisa yiligi rostmi:", x)

// }

// yettinchi masala Sonning bo‘luvchilari va uning turi
//bu funksiya sonning bo‘luvshilari nechtaligini topadi va uning turini aniqlaydi

func main() {
	var a = 72
	var b = 1
	var c = 8
	var d = 2
	var f = 36
	var x = 3
	var y = 24
	var s = 4
	var g = 18
	var h = 6
	var j = 12
	var k = 9
	fmt.Println("72 ning boluvshilari:", a, b, c, d, f, x, y, s, g, h, j, k)
	if a%a == 0 && a%b == 0 && a%c != 0 && a%d != 0 && a%f != 0 && a%x != 0 && a%y != 0 && a%s != 0 && a%g != 0 && a%h != 0 && a%j != 0 && a%k != 0 {
		fmt.Println("Tub:", a)
	} else if a == b+c+d+f+x+y+s+g+h+j+k {
		fmt.Println("Murakkab:", a)
	} else {
		fmt.Println("Oddiy son:", a)
	}

	sum := 0
	for i := 1; i < 72; i++ {
		if 72%i != 0 {
			continue
		}
		sum += i
	}
}
