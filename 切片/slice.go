package main
import "fmt"

func main() {
	// int 类型切片
	var slice []int
	fmt.Println(slice)

	var s1 []int
	var s2 []string
	s1 = []int{1,2,3}
	s2 = []string{"adsfasdf","asdfasdf","asdfasdf"}
	fmt.Println(s1,s2)
}