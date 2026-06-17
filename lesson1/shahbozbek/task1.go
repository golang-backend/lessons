// 1-masala: Slice bilan o'rtacha qiymat
// Butun sonlardan iborat slice bering: []int{12, 7, 25, 3, 18} .
// Slice ichidagi barcha sonlar yig'indisini hisoblang.
// O'rtacha ( average ) qiymatni toping.
// len() dan foydalaning, sonlar sonini qattiq (hardcoded) yozmang.
package main
import "fmt"
func average(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum = sum + num
	}
	return sum / len(slice)

}
// 2-masala: Slice'ni o'stirish
// Bo'sh slice ( var s []int ) yarating va append yordamida 1 dan 10 gacha
// sonlarni qo'shing. So'ng faqat juft sonlarni yangi slice'ga ajratib oling
// va chiqaring.
// Tekshiriladi: append , for sikli, % operatori, slice'ni filtrlash.
func A(nums []int) ([]int, []int) {
	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}
	var s []int
	for _, num := range nums {
		if num%2 == 0 {
			s = append(s, num)
		}
	}
	return nums, s
}
func main(){
	slice := []int{12, 7, 25, 3, 18}
	fmt.Println(average(slice))

	// masala 2

	
	var s []int
	fmt.Println(A(s))
}