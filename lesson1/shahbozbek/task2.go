// 3-masala: So'zlarni sanash
// Quyidagi slice berilgan:
// mevalar := []string{"olma", "anor", "olma", "uzum", "anor", "olma"}
// map[string]int yarating va har bir meva nechta marta uchraganini sanang.
// Natija: olma: 3, anor: 2, uzum: 1 .
package main  
import "fmt"
func caunts(mevalar []string) map[string]int {
	
	count := make(map[string]int)
	for _, meva := range mevalar {
		count[meva]++
	}
	return count

}
// 4-masala: Telefon kitobchasi
// map[string]string ko'rinishida telefon kitobchasi yarating (ism → raqam).
// Kamida 3 ta kontakt qo'shing.
// Berilgan ism bo'yicha raqamni qidiring va topilmasa "Topilmadi" deb
// chiqaring
// ( value, ok := m[key] shaklini ishlating).

func phone(contacts map[string]string, ism string) {

	value, ok := contacts[ism]
	if ok {
		fmt.Println("topildi", contacts[value], ok)
	} else {
		fmt.Println("topilmadi")
	}

}
func main(){
	mevalar := []string{"olma", "anor", "olma", "uzum", "anor", "olma"}
	fmt.Println(caunts(mevalar))
	// masala 4

	
	contacts := make(map[string]string)
	contacts = map[string]string{
		"Ali":  "+998901234567",
		"Vali": "+998937654321",
		"Olim": "+998991112233",
	}
	phone(contacts, "Olim")
}
