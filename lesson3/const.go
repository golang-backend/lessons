package main

import "fmt"

const (
	ism      string = "Temur"
	age      int    = 18
	phone    string = "+998-91-123-45-67"
	timeZone        = "Tashkent/Asia"
)

func main() {
	const (
		a = 10
		b
		c         = 234
		startTime = "09:00"
		endTime   = "18:00"
	)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

	// const (
	// 	dush = 0
	// 	sesh = 1
	// 	chor = 2
	// 	pay  = 3
	// 	juma = 4
	// 	shan = 5
	// 	yak  = 6
	// )

	// fmt.Println("dushanba:", dush)
	// fmt.Println("seshanba:", sesh)
	// fmt.Println("chorak:", chor)
	// fmt.Println("payshanba:", pay)
	// fmt.Println("juma:", juma)
	// fmt.Println("shanba:", shan)
	// fmt.Println("yakshanba:", yak)

	const (
		dush = iota
		sesh
		chor
		pay
		juma
		shan
		yak
		_
		add
	)

	fmt.Println("dushanba:", dush)
	fmt.Println("seshanba:", sesh)
	fmt.Println("chorak:", chor)
	fmt.Println("payshanba:", pay)
	fmt.Println("juma:", juma)
	fmt.Println("shanba:", shan)
	fmt.Println("yakshanba:", yak)
	fmt.Println("Add: ", add)
}
