// 5-masala: Talaba structi
// Student nomli struct yarating: Ism string , Yosh int , Baho float64 .
// 3 ta talaba yarating va ularni []Student slice'ga joylang.
// Slice bo'ylab aylanib ( for range ), bahosi 4.0 dan yuqori talabalarni
// chiqaring.
package main
import "fmt"
type Student struct {
	Ism  string
	Yosh int
	Baho int
}
// 6-masala: Method bilan struct
// 5-masaladagi Student ga method qo'shing:
// func (s Student) Otdimi() bool { ... }
// Agar baho 3.0 dan baland bo'lsa true qaytarsin. Har bir talaba uchun
// talaba.Otdimi() ni chaqirib natijani chiqaring.
// Tekshiriladi: method receiver, value vs pointer tushunchasi.
func (s Student) Otdimi() bool {

	if s.Baho > 3 {
		return true
	}

	return false
}




func main(){
	talabalar := []Student{
		{Ism: "Asadbek", Yosh: 20, Baho: 5.0},
		{Ism: "Madina", Yosh: 19, Baho: 4.0},
		{Ism: "Sardor", Yosh: 21, Baho: 3.0},
	}

	for _, talaba := range talabalar {
		if talaba.Baho >= 4.0 {
			fmt.Println(talaba.Ism)

		}

	}


	// masala 6
	talabalar = []Student{
		{Ism: "Asadbek", Yosh: 20, Baho: 5.0},
		{Ism: "Madina", Yosh: 19, Baho: 4.0},
		{Ism: "Sardor", Yosh: 21, Baho: 2.0},
	}
	for _, talaba := range talabalar {
		if talaba.Otdimi() {
			fmt.Println("talaba otdi", talaba.Ism)
		} else {
			fmt.Println("talaaba yiqildi", talaba.Ism)
		}
	}
	
}
