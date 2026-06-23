package main

import "fmt"

func main() {
	// var day uint8

	// fmt.Printf("Hafta kunini raqamda kiriting: ")
	// fmt.Scan(&day)
	// if else bilan
	// if day == 1 {
	// 	fmt.Println("Yakshanba")
	// } else if day == 2 {
	// 	fmt.Println("Dushanba")
	// } else if day == 3 {
	// 	fmt.Println("Seshanba")
	// } else if day == 4 {
	// 	fmt.Println("Chorshanba")
	// } else if day == 5 {
	// 	fmt.Println("Payshanba")
	// } else if day == 6 {
	// 	fmt.Println("Juma")
	// } else if day == 7 {
	// 	fmt.Println("Shanba")
	// } else {
	// 	fmt.Println("Noto'g'ri kun kiritingiz!")
	// }

	// switch case bilan
	// switch day {
	// case 1:
	// 	fmt.Println("Yakshanba")
	// case 2:
	// 	fmt.Println("Dushanba")
	// case 3:
	// 	fmt.Println("Chorshanba")
	// default:
	// 	fmt.Println("Noto'g'ri kun kiritingiz!")
	// }

	// switch case operator bilan
	// a := 34
	// switch {
	// case a > 0:
	// 	fmt.Println("a katta 0 dan ")
	// case a == 0:
	// 	fmt.Println("a 0 ga teng ")
	// default:
	// 	fmt.Println("a manfiy son")
	// }

	// har bir case ni majburan tekshirish
	var x = 5
	switch x {
	case 2:
		fmt.Println("2 ga teng")
		fallthrough
	case 5:
		fmt.Println("5 ga teng")
		fallthrough
	case 4:
		fmt.Println("4 ga teng")
		fallthrough
	case 6:
		fmt.Println("6 ga teng")
		fallthrough
	default:
		fmt.Println("topilmadi")
	}
}
