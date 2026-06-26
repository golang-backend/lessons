package main

import "fmt"

// 7 masala

func isPolindrom(soz string) {

	for i := 0; i < len(soz)/2; i++ {
		if soz[0] != soz[len(soz)-1-i] {  // salom -> s != m 2) s != o 
			fmt.Println("polindrom emas")
		}
		return
	}

	t := []rune(soz)

	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	if string(t) != soz {
		fmt.Println("Is Not Polindrom!!!")
	}
	fmt.Println("bu soz polindrom")
}

// 8- masala
func ortacha(sonlar []int) {
	caunt := 0
	sum := 0
	for _, num := range sonlar {
		if num < 0 {
			fmt.Println("bu son manfiy son")
			continue
		}
		if num > 0 {
			sum = sum + num
			caunt++

		}

	}
	average := sum / caunt
	fmt.Println(average)

}

func main() {
	// 7- masala
	isPolindrom("maktab")

	// 8- masala
	slice := []int{1, 2, 6, 8, 3, -6}
	ortacha(slice)

}
