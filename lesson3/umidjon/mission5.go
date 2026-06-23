package main
import "fmt"
// beshinchi masala Iota yordamida statuslat va solishtirish 

func main() {
const (
	StatusYuborildi = iota  // 0
	StatusQabulqilindi      // 1
	StatusRadetildi         // 2
)
    joriyStatus := StatusYuborildi   // bunda joriyStatus := o degani
	x := joriyStatus != StatusRadetildi  // 0 && 2 teng ekanligi yolg‘onmi
	fmt.Println(x)
}