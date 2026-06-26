package main
import "fmt"
// oltinchi masala Harorat chegaralari va operatoriy
// bu funksiyada hona harorati temp ozgaruvchiisi bilan e‘lon qilingan
// func main() {
// 	temp := 22  // bunda harorat 22 ligi aytilmoqda
// 	x := temp > 18 && temp < 26   // harorat 18 va 22 orag‘idaligi
// 	y := temp == 18 || temp == 26 // harorat 18 yoki 22 ga tengligi rostmi
// 	fmt.Println(x)
// 	fmt.Println(y) 
// }


// oltinchi masala O‘quvchilarni baholash
// bu funksiya o‘quvchilarni to‘plagan baliga qarab ularga baho qoyadi

func main() {
	var temur = 90
	var bobur = 72
	var nodir = 35
	if temur > bobur && temur > nodir {
		fmt.Println("winner is Temur")
	}
	if bobur > temur && bobur > nodir {
		fmt.Println("Winner is Bobur")
	}
	if nodir > temur && nodir > bobur {
		fmt.Println("Winner is Nodir")
	}
	switch {
	case temur >= 85 && temur <= 100:
	    fmt.Println("Temur 5 baho oldi")
		fallthrough
	case bobur >= 70 && bobur <= 84:
		fmt.Println("Bobur 4 baho oldi")
	    fallthrough
	case nodir >= 55 && nodir <= 69:
		fmt.Println("Nodir 3 baho oldi")
		fallthrough
	default:
		fmt.Println("Boshqa o‘quvchi yo‘q")
	}
}