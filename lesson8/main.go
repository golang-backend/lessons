package main

import "fmt"

func main() {
	// declaration:
	// var slice = []int{1, 2, 4, 3, 2}

	// slice1 := []int{}

	slice2 := make([]int64, 10, 20) // length = 10, capacity = 20

	fmt.Println("slice2: ", slice2)

	a := []int8{11, 3, 45, -1, 3, 8, -2}

	baseSlice := a[0:3]

	baseSlice1 := a[4:]

	fmt.Println(a, baseSlice)
	fmt.Println(baseSlice1)

	// phoneArray := make([]string, 100)

	// phone := ""
	// fmt.Scan(&phone)
	// for i := 0; i < len(phoneArray); i++ {
	// 	phoneArray[i] = phone
	// }

	// phone2 := "+998911234567"

	// phoneArray = append(phoneArray, phone2)

	k := []int{1, 2, 3, 5}

	fmt.Printf("Length: %d\nCapacity: %d\n", len(k), cap(k))

	k = append(k, []int{2, 3, 4}...)

	fmt.Printf("Length: %d\nCapacity: %d\n", len(k), cap(k))

	m := []int{1, 2, 5}
	n := make([]int, 3)

	fmt.Println("m = ", m)
	fmt.Println("n = ", n)

	copy(n, m)

	fmt.Println("m = ", m)
	fmt.Println("n = ", n)

	numbers := []string{"123", "45", "90", "-10", "34", "23", "78"}
	fmt.Println("Numbers: ", numbers)
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("Numbers2: ", numbers)

}
