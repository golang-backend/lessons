package main

import "fmt"

func main() {
	// var a, b int
	// fmt.Printf("Enter two numbers: ")
	// fmt.Scan(&a, &b)

	// max(a, b)
	// arr := []int{3, 45, 6, 2, 7, 1}
	// min(arr)

	g := sum(34, 05, 20)
	fmt.Println("Sum: ", g)

	name, age := userInfo("John", "Doe", 30)

	fmt.Println("Name: ", name, "Age: ", age)

	// _, msg := checking(20)
	// fmt.Println("Message: ", msg)

	// univer(42)

	// universal(1, 2, 3, "Hello", "World", true, 3.14)

	add(10, 20, func() {
		fmt.Println("Adding two numbers...")
	})

}

// max function takes two integers and prints the maximum value or indicates if they are equal.
func get(x int, y int) {
	if x > y {
		fmt.Println("Max: ", x)
	} else if y > x {
		fmt.Println("Max: ", y)
	} else {
		fmt.Println("Equal")
	}
}

// min function takes a slice of integers and prints the minimum value.
func min(nums []int) {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	fmt.Println("Min: ", min)
}

func sum(x, y, z int) int {
	s := x + y + z
	return s
}

func userInfo(firstName string, lastName string, age int) (string, int) {
	user := firstName + " " + lastName
	return user, age
}

func checking(h int) (int, string) {
	if h > 25 {
		return h, "You are above 25"
	}
	return h, "You are below 25"
}

func univer(b any) int {
	switch v := b.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	case bool:
		fmt.Println("Boolean: ", v)
	default:
		fmt.Println("Unknown type")
	}

	return 0
}

func universal(a ...any) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func add(x, y int, print func()) int {
	print()
	j := func(sum int) {
		fmt.Println("SUM: ", sum)
	}

	fmt.Println("--->> ", j)

	j(x + y)
	
	return x + y
}

func printAdd(sum int) {
	fmt.Println("SUM: ", sum)
}
