package main
import "fmt"

// beshinchi masala  Yil fasli va harorat tavsiyasi
// bu funcsiya faslni aniqlab qanday kiyim kiyishni aytadi

func main() {
	var month uint
	fmt.Scan(&month)
	fmt.Println("Month:", month)
	if month == 12 || month == 1 || month == 2 {
		fmt.Println("Winter: wear warm clothes") // Qish: issiq kiyim kiying
	}
	if month == 3 || month == 4 || month == 5 {  // Bahor: yengil kiyim kiying
		fmt.Println("Spring: wear light clothes")
	}
	if month == 6 || month == 7 || month == 8 {  // Yoz: futbolka kiying
		fmt.Println("Summer: wear T-shirts")
	}
	if month == 9 || month == 10 || month == 11 {
		fmt.Println("Autumn: wear light clothes")  // Kuz: yengil kiyim kiying
	}
}


