package main


var x,y int
var ( // 一般用于声明全局变量
	a int
	b bool
)

var c,d int = 1,2
var e,f = 123,"hello"

func  main()  {
	g,h := 12323434,"sdfadf" // 不带声明格式的只能在函数体中出现
	println(x,y,a,b,c,d,e,f,g,h)
}


