package main

import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main()  {
	fmt.Println(Books{"Go语言","www.baidu.com","语言教程",2000})

	// 忽略字段
	fmt.Println(Books{title:"go",author:"www.baidu.com"})

	// 访问结构体成员

	var Book1 Books
	var Book2 Books
	Book1.title = "go"
	Book1.author = "www.baidu.com"
	Book1.subject = "hehe"
	Book1.book_id = 2123123

	Book2.title = "oc"
	Book2.author = "www.oc.com"
	Book2.subject = "objective-c"
	Book2.book_id = 2000

	fmt.Printf("Book1 title :%s \n",Book1.title)
	fmt.Printf("Book1 author :%s \n",Book1.author)
	fmt.Printf("Book1 subject :%s \n",Book1.subject)
	fmt.Printf("Book1 book_id :%d \n",Book1.book_id)

	fmt.Printf("Book2 title :%s \n",Book2.title)
	fmt.Printf("Book2 author :%s \n",Book2.author)
	fmt.Printf("Book2 subject :%s \n",Book2.subject)
	fmt.Printf("Book2 book_id :%d \n",Book2.book_id)





}
