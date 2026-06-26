package main

import "fmt"

func main() {
	// masala 3
	nums := []int{1, 3, 3, 1, 0}
	max := 0
	count := 0
	for _, num := range nums {
		if num > max {
			max = num
			count = 1

		} else if num == max {
			count++

		}

	}
	fmt.Println("katta son", max)
	fmt.Println("takrorlangan", count)
	// masala 5

	numss := []int{30, 11, 7, 101}

	for _, num := range numss {

		if num < 10 {
			continue
		}

		if num > 100 {
			break
		}

		fmt.Println(num)
	}
}
