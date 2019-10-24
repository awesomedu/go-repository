package main
import "fmt"

// 求和
func main(){
	a1 := [...]int{1,3,5,6,7}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)
}