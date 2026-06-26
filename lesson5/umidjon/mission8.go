package main
import "fmt"
// sakkizinchi masala Chegirma olish huquqi (yoki operatoriy)


// func main() {
// 	const IsVip = false
// 	summa := 600000
// 	a := IsVip || summa > 500000
// 	fmt.Println("Mijoz chegirma oladimi:", a)
// }


// sakkizinchi masala Koordinata tekisligi

func main() {
	var (
		x1 int
		x2 int
		x3 int
		y1 int
		y2 int
		y3 int
		
	)
	fmt.Scan(&x1, &x2, &x3, &y1, &y2, &y3)
	fmt.Println(x1, x2, x3, y1, y2, y3)
	fmt.Println("A:", x1, y1)
	fmt.Println("B:", x2, y2)
	fmt.Println("C:", x3, y3)
	if (x1 - x2)*(y1 - y2) == (x1 - x3) * (y1 - y3) {
		fmt.Println("Nuqtalar birlashtirilsa to‘g‘ri chiziq hosil bo‘ladi:")
		return
    }
	var x = (x1 - x2)*(x1 - x2) + (y1 - y2)*(y1 - y2)
	var y = (x2 - x3)*(x2 - x3) + (y2 - y3)*(y2 - y3)
	var z = (x1 - x3)*(x1 - x3) + (y1 - y3)*(y1 - y3)
	switch {
	case x > y + z || y > x + z || z > x + y:
		fmt.Println("O‘tmas burchakli uchburvhak hosil bo‘ldi")
    case x < y + z || y < x + z || z < x + y:
		fmt.Println("O‘tkir burchakli uchburchak hosil bo‘ldi")
	case x == y + z || y == x + z || z == x + y:
		fmt.Println("To‘g‘ri burchakli uchburchak hosil bo‘ldi")
    }
}

