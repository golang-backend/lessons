package main

import "fmt"

func main() {
	// Array Declaration:
	// var massiv = [3]int{2, 9, 7}
	// //index:    0  1  2
	// mass := [5]float64{3.14, 2.71, 6.667, 19.23, 1.2}
	// // Get Array Length:
	// massiv1Length := len(massiv)
	// mass1Length := len(mass)
	// fmt.Println("M1: ", massiv1Length, "M2: ", mass1Length)
	// // Array declaration without number:
	// arr := [...]string{"olma", "anor", "uzum", "tarvuz", "olcha", "qovun", "nok"}
	// fmt.Println("Arr Length: ", len(arr))
	// a := [4]int{}
	// fmt.Println("a: ", a)

	// a[0] = 3

	// fmt.Println("a: ", a)

	// // Array declaration with indexing element:
	// arr1 := [5]int{1: 34, 4: -5}

	// fmt.Println("arr1: ", arr1)

	var arr = [8]int{67, 2, 3, 5, 10, 3, 0, -1}

	for i := 0; i < len(arr); i++ { // 0 , 1, 2, 3,
		// fmt.Printf("index: %d: %d\n", i, arr[i])
		// if arr[i] < 0 {
		// 	arr[i] = arr[i] * (-1)
		// }

		if arr[i]%2 == 0 { // 1) a[i] = a[0] = 67
			// 2) a[i] = a[1] = 2
			// fmt.Println("Pair: ", arr[i])
		}
	}
	count := 0
	for i, v := range arr {
		fmt.Println("index: ", i, "value: ", v)
		if v == 3 {
			count++
		}
	}
	
	fmt.Println(count)
}
