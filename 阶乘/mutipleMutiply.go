package main

import "fmt"

func count(n uint64) (result uint64) {
	if n > 0 {
		result = n * count(n - 1)
		return result
	}else {
		return 1
	}
}

func main()  {
	var  i int = 20
	fmt.Printf("%d 的阶乘是 %d\n",i,count(uint64((i))))
}
