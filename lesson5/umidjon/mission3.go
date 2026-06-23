package main

import "fmt"


//  uchinchi masala Ro‘yxatdan o‘tish yoshi(constants)
//  bu funksiya foydalanuvchini royxatdan otish || otmasligini aniqlaydi
// func main() {
// 	const minAge = 18 // ro‘yxatdan o‘tish yoshi
// 	userAge := 16  // foydalanuvshi yoshi
// 	result := userAge >= minAge
// 	fmt.Println("Foydalanuvchi ro‘yxatdan o‘tgani rostmi", result)
// 	// natijada foydalanuvch ro‘yxatdan o‘tgan bo‘lsa true o‘tmagan bolsa false chiqaradi
// }

// uchinchi masala O‘quvchi bahosi
// bu funksiya oquvchining baliga qarab darajasini aniqlaydi
// func main() {
// 	var ball int
// 	fmt.Scan(&ball)
// 	fmt.Println("Ball:", ball)
// 	if ball >= 90 && ball <= 100 {
// 		fmt.Printf("Excelent: %d\n", ball) // Excelent - a‘lo
// 	}
// 	if ball >= 70 && ball <= 89 {
// 		fmt.Printf("Good: %d\n", ball) // good - yaxshi
// 	}
// 	if ball >= 50 && ball <= 69 {
// 		fmt.Printf("Satisfactory: %d\n", ball) // satisfactory - qoniqarli
// 	}
// 	if ball >= 0 && ball <= 49 {
// 		fmt.Printf("Unsatisfactory: %d\n", ball) // unsatisfactory - qoniqarsiz
// 	}
//     if ball > 100 || ball < 0 {
// 		fmt.Printf("Wrong ball: %d\n", ball)  // wrong ball - ball xato kiritildi
// 	}
// }

// uchinchi masala Eng katta, eng kichik && o‘rtancha qiymat
// bu funksiya quyidagi 5 ta sonning eng katta, eng kichik && o‘rtancha qiymatlarini topadi

func main() {
	a := 27
	b := 24
	c := 34
	d := 12
	f := 21
	fmt.Println("A:", a, "B:", b, "C:", c, "D:", d, "F:", f)
    if a < b {
		a = b 
	}
	if a < c {
		a = c
	}
	if a < d {
		a = d
	}
	if a < f {
		a = f
	}
	fmt.Println("Big number:", a)
	switch {
	case a > b:
		a = b
		fallthrough
	case a > d:
		a = d
	case a > c:
		a = c
	case a > f:
		a = f
	}
	fmt.Println("Small number:", a)
    
}
