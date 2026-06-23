package main
import "fmt"
// ikkinchi masala Dokon balansi && narx
// bu funksiya balans maxsulotni sotib olishga yetarli || yetarli emasligini aniqlaydi 
func main() {
	balance := 45000  //  45000 mijozning balansi
	charge := 50000  // 50000 mahsulot narxi
	result := balance >= charge 
	fmt.Println("Result: ",result)
}