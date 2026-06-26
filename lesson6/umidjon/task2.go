package main
import "fmt"
// bu funksiya 1 dan 50 gacha bolgan juft sonlar kopaytmasini topadi
func main() {
	var sum int64 = 1
	for i := int64(1); i <= 50; i++ {  
		if i%2 != 0 {
			continue
		}
		sum *= i
    }
	fmt.Println("kopaytma:", sum)
}