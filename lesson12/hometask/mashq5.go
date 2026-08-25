package main
import "fmt"
func main() {
    var txt string
    fmt.Scan(&txt)
    r := []rune(txt)
    j := len(r)-1
    for i := 0; i < j; i++ {
        if r[i] != r[j] {
            fmt.Println("yuq")
            return
        } 
        j--
    }
    fmt.Println("Palindrom")
}