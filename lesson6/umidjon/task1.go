package main
import "fmt"
func main() {
	var n int
	var sum = 0
	fmt.Scan(&n)
	fmt.Println("N:", n)
	for i := 0; i <= int(n); i++ {
		sum += i
    }
	fmt. Println("Yig‘indi", sum)
}