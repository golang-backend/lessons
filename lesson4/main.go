package main

import "fmt"

func main() {
	var x int
	var y int
	// 1-shart x soni 0 dan katta bo'lsa y soniga qo'shing
	fmt.Printf("x sonni kiriting: ")
	fmt.Scan(&x) // & - ampersand bilan yozilsa x o'zgaruvchiga terminaldan olinga qiymatni beradi
	fmt.Printf("y sonni kiriting: ")
	fmt.Scan(&y)

	if x > 0 {
		// shart bajarilsa shu figurali qavsni ichida mantiqni davom ettiramiz
		fmt.Println("x + y = ", x+y)
	} else if x < 0 && y > 0 {
		fmt.Println(" x * y = ", x*y)
	} else {
		fmt.Println("x = ", x)
	}

	// 2-shart y soni ham 0 katta bo'lsagina ikki sonni yig'indishi ekranga chop eting
	if x > 0 && y > 0 {
		fmt.Println("2: x + y = ", x+y)
	}

}
