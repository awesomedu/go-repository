package main

import "fmt"

var a,b int
var c string




func main()  {
	a,b,c = 5,7,"string2" // 这种情况下不报错
	fmt.Println(a,b)

	a,b:= 5,123 // 这种情况会报  c declared and not used
	fmt.Println(a,b)
}
