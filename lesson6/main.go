package main

import "fmt"

func main() {
	// var n = 5

	// i := 1

	// if i <= n {
	// 	i = i + 1
	// }
	// fmt.Println("i = ", i)

	// for i < n { // 1) 1 < 5 = true 2) 2 < 5       3) 3 < 5
	// 	i = i + 1 // 1 + 1 = 2     2) 2 + 1 = 3   3) 3 + 1 = 4
	// }

	// fmt.Println("i = ", i)

	k := 1

	for k <= 10 {
		if k%2 == 0 {
			// fmt.Println("=> ", k)
		}
		k++ // k = k + 1
	}

	for i := 1; i < 5; i++ {
		if i == 2 {
			fmt.Println("*")
			break
		}
	}

	a := "s"

	var word string
	for i := 0; i < 10; i++ {
		word = word + a
	}
	// fmt.Println("-> ", word)

	s := "salom" // 0 index -> s 1-index = a
	h := ""
	for i := range s {

		h += string(s[i])

		if string(s[i]) == "l" {
			fmt.Println("= ", string(s[i]))
			break
		}
	}
	fmt.Println("Partial text: ", h)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println("-> ", string(s[i]))
	// }

	// for {
	// 	fmt.Println("Cheksiz")
	// }
	var sum = 0
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i = ", i)
		sum += i
	}
	fmt.Println("SUM: ", sum)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 || i != 5 || j == 0 || j != 5 {
				fmt.Println("* ")
			}
		}
		fmt.Println()
	}
}
