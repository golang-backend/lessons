package main

import "fmt"

func engUzun(slice []int) {
	if len(slice) == 0 {
		fmt.Println(0)
		return
	}

	current := 1
	maxLen := 1

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			current++
		} else {
			if current > maxLen {
				maxLen = current
			}
			current = 1
		}
	}

	// oxirgi ketma-ketlikni ham tekshirib qo'yamiz
	if current > maxLen {
		maxLen = current
	}

	fmt.Println(maxLen)
}
func main() {
	arr := []int{1, 2, 3, 2, 4, 5, 6, 1}
	engUzun(arr)

}
