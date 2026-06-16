package main

import "fmt"

type User struct {
	Name string
	Age  int
}


func main() {
	fmt.Println("Hello, World!")
	a := User{Name: "Alice", Age: 30}
	fmt.Printf("User: %+v\n", a)

	v := Person{Name: "Bob", Age: 25}
	fmt.Printf("Person: %+v\n", v)
}

type Person struct {
	Name string
	Age  int
}