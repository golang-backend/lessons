package main

import "fmt"

func main() {
	// 1 masala
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum = sum + num

	}
	fmt.Println(sum)

	// 2 masala

   n := []int{38, 24, 800, 8, 16}

	sum2 := 0

	for i := 0; i < len(n); i++ {
		num := n[i]

		if num >= 10 && num <= 99 && num%8 == 0 {
			sum2 += num
		}
	}

	fmt.Println(sum2)
}

