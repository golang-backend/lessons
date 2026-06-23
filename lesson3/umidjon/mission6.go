package main

import "fmt"

// oltinchi masala Harorat chegaralari va operatoriy
// bu funksiyada hona harorati temp ozgaruvchiisi bilan e‘lon qilingan
func main() {
	temp := 22                    // bunda harorat 22 ligi aytilmoqda
	x := temp > 18 && temp < 26   // harorat 18 va 22 orag‘idaligi
	y := temp == 18 || temp == 26 // harorat 18 yoki 22 ga tengligi rostmi
	fmt.Println(x)
	fmt.Println(y)
}
