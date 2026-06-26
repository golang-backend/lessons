package main

import (
	"fmt"
)

// to‘qqizinchi masala Uvhburchak sharti
// bu funksiya uchburchak || uchburchak emasligini aniqlaydi

// func main() {
// 	a, b, c := 5, 8, 10
// 	 v  := a + b > c && a +c > b && b + c > a
// 	fmt.Println(v)
// }

// toqqizinchi masala 3 xonali son
// bu funksiya 3 xonali sonni raqamlari ustida amallar bajaradi

func main() {
	x := 325
	fmt.Println("X:", x)
	a := x % 10
	b := (x / 10) % 10
	c := x / 100
	if (a+b+c)%2 == 0 {
		fmt.Println("Pair")
	}
	if (a+b+c)%2 != 0 {
		fmt.Println("Odd")
	}
	if a > b && a > c {
		fmt.Println("big number:", a)
		fmt.Println("Razriyadi: birlar")
	} else if b > a && b > c {
		fmt.Println("big number:", b)
	} else {
		fmt.Println("big number:", c)
	}
	if x == a*a*a+b*b*b+c*c*c {
		fmt.Println("Armstrong number:", x)
	}

}
