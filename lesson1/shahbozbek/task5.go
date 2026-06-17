// 10-masala: Mini do'kon tizimi
// Quyidagilarni birlashtiring:
// Mahsulot structi: Nomi string , Narxi float64 , Soni int .
// []Mahsulot slice — kamida 4 ta mahsulot.
// map[string]float64 — mahsulot nomi → umumiy qiymati ( Narxi * Soni ).
// Eng qimmat mahsulotni topadigan funksiya yozing.
// Qimmatbaho() bool methodi: narxi 100000 dan oshsa true .
package main 
import "fmt"
type Mahsulot struct {
	Nomi  string
	Narxi float64
	Soni  int
}

func mostExpensive(mahsulotlar []Mahsulot) Mahsulot {
	if len(mahsulotlar) == 0 {
		return Mahsulot{} 
	}
	expensive := mahsulotlar[0]
	for _, mahsulot := range mahsulotlar {
		if mahsulot.Narxi > expensive.Narxi {
			expensive = mahsulot
			fmt.Println("eng qimmat mahsulot: ", mahsulot.Nomi, expensive)

		}
	}
	return expensive

}
func (mahsulotlar Mahsulot) Qimmatbaho() bool {
	return mahsulotlar.Narxi > 100000
}
func main(){
	
	mahsulotlar := []Mahsulot{
		{Nomi: "Noutbuk", Narxi: 8500000.0, Soni: 12},
		{Nomi: "Telefon", Narxi: 4200000.0, Soni: 25},
		{Nomi: "Monitor", Narxi: 1800000.0, Soni: 8},
		{Nomi: "Klaviatura", Narxi: 250000.0, Soni: 50},
	}
	for _, mahsulot := range mahsulotlar {
		fmt.Println("mahsulot nomi->", mahsulot.Nomi, "umumiy qiymati=", mahsulot.Soni*int(mahsulot.Narxi))
	}
	mostExpensive(mahsulotlar)
	for _, mahsulot := range mahsulotlar {
		fmt.Println(mahsulot.Qimmatbaho(), mahsulot.Nomi)
	}
}