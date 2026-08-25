package main

import "fmt"

func main() {
	// map -> kalit va qiymatlar juftligini saqlaydi
	// var dic map[string]int // noto'g'ri e'lon qilish
	// dic["ds"] = 12
	// fmt.Println(dic)

	// var n = make(map[string]int) // make yordamida e'lon qilish
	// n["apple"] = 15

	// fmt.Println("n: ", n)

	// kalit va qiymatlarni biriktirish orqali e'lon qilish
	// dictionary := map[int]int8{
	// 	123: 34,
	// 	345: 12,
	// 	124: 87,
	// }

	// fmt.Println("-->> ", dictionary)

	// fmt.Println("x: ", dictionary[123])

	// delete(dictionary, 124)

	// fmt.Println(dictionary)

	// mevalar := map[string]float64{
	// 	"apple":  10500,
	// 	"banana": 30000,
	// 	"orange": 12000,
	// }

	// // price, isExists := mevalar["nok"]
	// if price, ok := mevalar["banana"]; ok {
	// 	fmt.Println("Price: ", price)
	// }

	// for key, value := range mevalar {
	// 	if key == "apple" {
	// 		mevalar[key] = mevalar[key] + 1000
	// 	}
	// 	fmt.Printf("%s: %.2f\n", key, value)
	// }
	// fmt.Println("-->> ", mevalar)

	users := map[int]User{
		1: User{
			name:     "Ali",
			age:      20,
			phone:    "+998911234567",
			birthday: "2000-06-19",
		},
		2: User{
			name:     "Vali",
			age:      22,
			phone:    "+998911238976",
			birthday: "2005-06-10",
		},
	}
	users = users

	var arr = []string{"olma", "banan", "nok", "olma", "nok", "uzum", "banan", "nok", "uzum"}

	fruits := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		fruits[arr[i]]++
	}

	fmt.Println("Fruits: ", fruits)
	fmt.Println("Length: ", len(fruits))

}

type User struct {
	name     string
	age      int
	phone    string
	birthday string
}
