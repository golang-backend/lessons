package main

import "fmt"

// birinchi masala massivning elementlri kvadrati
func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	fmt.Println("kopaytma:", arr)

	// ikkinchi masala massivning eng katta elementini aniqlash
	arr1 := [...]int{2, -15, 65, 89, -578, 61, 120}

	if arr1[0] < arr1[1] {
		arr1[0] = arr1[1]
	}
	if arr1[0] < arr1[2] {
		arr1[0] = arr1[2]
	}
	if arr1[0] < arr1[3] {
		arr1[0] = arr1[3]
	}
	if arr1[0] < arr1[4] {
		arr1[0] = arr1[4]
	}
	if arr1[0] < arr1[5] {
		arr1[0] = arr1[5]
	}
	if arr1[0] < arr1[6] {
		arr1[0] = arr1[6]
	}
	fmt.Println("Eng katta element:", arr1[0])

	// uchinchi masala
	var n1 int
	fmt.Printf("N sonni kirit: ")
	fmt.Scan(&n1)
	a := n1
	arr3 := [a]int{}
	for i := 0; i < n1; i++ {
		fmt.Printf("elementni kirit: ")
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr3)

	// tortinchi masala eng katta va eng kichik elementlar

	var min = mass[0]
	var max = mass[0]
	for f := 0; f < len(mass); f++ {
		if max < mass[f] {
			max = mass[f]
		}
		if min > mass[f] {
			min = mass[f]
		}
	}
	fmt.Println("Max: ", max)
	fmt.Printf("Min:  %d", min)
}
