package main

import (
	"fmt"
	"lesson9/token"
)

func main() {
	// sh.BuildShape()
	// Println("Shape Built!!!")
	token.Init()
	fmt.Println(token.Generate())
	fmt.Println(token.Generate())
	fmt.Println(token.Generate())
}
