package main

import "fmt"

func main()  {
	var a1[3]bool
	var a2[4]bool
	fmt.Println(a1,a2)
	a1 = [3]bool{true,false,true}
	fmt.Println(a1)

	/// ... 根据初始值推测数组长度
	a100 := [...]int{1,2,2,123,1,23123,123,123}
	fmt.Println(a100)

	a3 := [5]int{0:3,4:2}
	fmt.Println(a3)

	/// 遍历
	citys := [...]string{"北京","上海","深圳"}
	// 索引
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
		fmt.Println(citys[1])
	}
	// for range
	for i, v := range citys {
		fmt.Println(i,v)
	}

	// 多维数组 [[1,2],[3,4],[5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println(a11)

	// 多维数组遍历
	for _, v1 := range a11 {
		// fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

}