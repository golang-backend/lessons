package main

import "fmt"

// 5-masala
func jufttoq(n int) {
	if n%2 == 0 {
		fmt.Println(n, "-> juft son")
	} else {
		fmt.Println(n, "-> toq son")
	}

}

func main() {
	// 4- raqamlar yigindisi    masala
	sum := 0
	num := 2589
	for num > 0 {
		x := num % 10
		sum = sum + x
		num = num / 10

	}
	fmt.Println(sum)

	// 5 masala
	jufttoq(1)

	// 6- masala
	for i := 1; i <= 10; i++ {
		for b := 1; b <= 10; b++ {
			fmt.Printf("%d x %d = %d\n", i, b, i*b)
		}
		fmt.Println()
	}

}
