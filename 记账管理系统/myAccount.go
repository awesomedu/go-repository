package main
import "fmt"

func main(){
    key:=""
    loop := true
    // 定义账户余额
    balance := 10000.0
    // 每月收支金额
    money := 0.0
    // 每次收支说明
    note := ""
    // 收支详情字符串记录
    details := "收支\t账户余额\t收支金额\t说明"

    for{
    // 接收输入
    fmt.Println("\n------------记账软件-----------")
    fmt.Println("------------1.收支明细-----------")
    fmt.Println("------------2.登记收入-----------")
    fmt.Println("------------3.登记支出-----------")
    fmt.Println("------------4.退出软件-----------")
    fmt.Println("------------请选择1~4-----------")
    fmt.Scanln(&key)
    switch key{
    	case "1":
    	fmt.Println("------------当前收支明细------------")
    	fmt.Println(details)
    	case "2":
    	fmt.Println("------------本次收入金额------------")
    	fmt.Scanln(&money)
    	balance += money
    	fmt.Println("------------本次收入说明-----------")
    	fmt.Scanln(&note)
    	fmt.Println("------------登记支出------------")
    	details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
    	case "3":
    	fmt.Println("------------当前收支明细------------")
    	case "4":
    	loop = false
    	default:
    	fmt.Println("------------请输入正确的选项------------")
    }
   		if !loop{
    		break
   		}
	}
}
