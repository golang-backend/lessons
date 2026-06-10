package main

import "fmt"

// 1-dars: O'zgaruvchilar va ma'lumot turlari

// Birinchi dastur
// func main() {
// 	// declare variables
// 	x := 2 // butun
// 	y := 12 // butun
// 	z := x + y
// 	// kars = 3.14
// 	var ism = "Umidjon" // text -> string
// 	fmt.Println(z)
// 	fmt.Println(ism)
// }

// Ma'lumot turlari
func main() {
	// integer - butun son
	var x int = 15 // butun son // int = 4 byte -> 32 bit // define by computer CPU
	// int8 = -128 dan 127 gacha -> -2^7 dan 2^7-1 gacha
	var y int8 = 127 // int8 = 1 byte

	var z int16 = -32767 // int16(16 bit) = 2 byte -2^15 dan 2^15-1 gacha

	var a int32 = 2147483647 // int32(32 bit) = 4 byte -2^31 dan 2^31-1 gacha

	var b int64 = 9223372036854775807 // int64(64 bit) = 8 byte -2^63 dan 2^63-1 gacha

	var c uint = 15 // unsigned integer - musbat butun son

	var d uint8 = 255 // unsigned integer 8 bit - 0 dan 255 gacha 2^8-1 gacha

	var e uint16 = 65535 // unsigned integer 16 bit - 0 dan 65535 gacha 2^16-1 gacha

	var f uint32 = 4294967295 // unsigned integer 32 bit - 0 dan 4294967295 gacha 2^32-1 gacha

	var g uint64 = 18446744073709551615 // unsigned integer 64 bit - 0 dan 18446744073709551615 gacha 2^64-1 gacha

	// float - o'nlik son
	var h float32 = 3.14              // float32 - 4 byte - 6-7 ta raqamни saqlaydi
	var i float64 = 3.141592653589793 // float64 - 8 byte - 15-16 ta raqamни saqlaydi

	// string - matn
	var ism string = "Umidjon go dasturlash tilini o'rganmoqda" // string - matn

	// 1 > 0 mi? rost -> tog'ri -> true
	// -1 > 0 mi? yolg'on -> noto'g'ri -> false

	// boolean - mantiqiy qiymat
	var isTrue bool = true   // rost -> tog'ri
	var isFalse bool = false // yolg'on -> noto'g'ri

	// x = 230
	// agar x 230 bo'lsa bu rost, agar x 230 bo'lmasa bu yolg'on
	// x == 230 -> true, x != 230 -> false

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(ism)
	fmt.Println(isTrue)
	fmt.Println(isFalse)

}
