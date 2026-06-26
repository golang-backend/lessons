package main

import (
	"fmt"
	"strconv"
)

// O‘ninchi masala Vaqt

func main() {
	var time1, time2 string
	fmt.Printf("Birinchi vaqt: ")
	fmt.Scan(&time1)
	fmt.Printf("Ikkinchi vaqt: ")
	fmt.Scan(&time2)

	fmt.Println("Times: ", time1, time2)

	hour1Str, minute1Str := time1[0:2], time1[3:5]
	hour2Str, minute2Str := time2[0:2], time2[3:5]

	hour1, _ := strconv.Atoi(hour1Str)
	minute1, _ := strconv.Atoi(minute1Str)

	hour2, _ := strconv.Atoi(hour2Str)
	minute2, _ := strconv.Atoi(minute2Str)

	


	fmt.Println("Hour1: ", hour1Str, minute1Str)

}
