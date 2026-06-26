package main
import "fmt"
// 


func main() {
	var a string
	var b string
	fmt.Println("son yoki soz kiriting")
	fmt.Scan(&a)
	for i := len(a)-1; i >= 0; i-- {
		b += string(a[i]) 
	}
	fmt.Println("B:", b)
	if a == b {
		fmt.Println("palindrom", a )
	} else {
		fmt.Println("palindrom emas", a)
	}
}