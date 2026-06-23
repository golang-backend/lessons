package main

import "fmt"

//  uchinchi masala Ro‘yxatdan o‘tish yoshi(constants)
//  bu funksiya foydalanuvchini royxatdan otish || otmasligini aniqlaydi
func main() {
	const minAge = 18 // ro‘yxatdan o‘tish yoshi
	var userAge = 0   // foydalanuvshi yoshi
	fmt.Printf("Yoshingizni kiriting: ")
	fmt.Scan(&userAge)
	result := userAge >= minAge
	fmt.Println("Foydalanuvchi ro‘yxatdan o‘tgani rostmi", result)
	// natijada foydalanuvch ro‘yxatdan o‘tgan bo‘lsa true o‘tmagan bolsa false chiqaradi
}
