// 7-masala: Ichma-ich (nested) struct
// Manzil structi yarating ( Shahar , Kocha ). Keyin Student ichiga
// Manzil ni maydon sifatida joylang. Bitta talaba yaratib, uning shahrini
// talaba.Manzil.Shahar orqali chiqaring.
package main
import "fmt"

type manzil struct {
	shahar string
	kocha  string
}
type student struct {
	talaba string
	manzil manzil
}
// 8-masala: Shakl yuzasi (klassik)
// Shape (Shakl) interfeysini yarating:type Shape interface {
// Yuza() float64
// }
// Doira (radius bilan) va Tortburchak (eni va bo'yi bilan) structlarini
// yarating.
// Ikkalasi ham Yuza() methodini amalga oshirsin.
// []Shape slice yaratib, har bir shaklning yuzasini bitta for sikli ichida
// chiqaring.
// Bu masala interfeysni chinakam tushunganini ko'rsatadi — har xil
// structlarni
// bitta interface orqali bir xil ishlata olish.
type Shape interface {
	Yuza() float64
}
type doira struct {
	radius float64
}
type tortburchak struct {
	eni  int
	boyi int
}

func (d doira) Yuza() float64 {
	return 3.14 * d.radius * d.radius
}
func (t tortburchak) Yuza() float64 {
	return float64(t.eni * t.boyi)
}

func main(){
	Talaba := student{
		talaba: "shahboz",
		manzil: manzil{
			shahar: "anadijon",
			kocha:  "posponlar",
		},
	}
	fmt.Println("kochasi->", Talaba.manzil.kocha)


	// masala 8
	shapes := []Shape{
		doira{radius: 5},
		tortburchak{4, 6},
		doira{radius: 3},
		tortburchak{10, 2},
	}
	for i, shape := range shapes {
		fmt.Printf("%d-shakl yuzasi: %.2f\n", i+1, shape.Yuza())
	}
}

