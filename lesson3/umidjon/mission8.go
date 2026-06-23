package main

import "fmt"

// sakkizinchi masala Chegirma olish huquqi (yoki operatoriy)

func main() {
	const IsVIP = false
	summa := 200000.0
	var chegirma = IsVIP || summa > 500000
	fmt.Println("Chegirma berilsinmi: ", chegirma)

	const IsBlocked = false

	var IsAccess = !IsBlocked // -->> noto'g'riligiga tekshiradi

	fmt.Println("Mijoz kira oladimi: ", IsAccess)
}
