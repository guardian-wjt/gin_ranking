package main

import (
	"gin_ranking/router"
)

func main() {
	//路由文件封装
	r := router.Router()

	//异常捕获
	// defer recover panic nil
	//defer func() {
	//	if err := recover(); err != nil { //recover	回复程序
	//		fmt.Println("捕获异常：", err)
	//	}
	//}()

	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3) //倒叙执行
	//panic("111")         //中断程序

	r.Run(":9999")
}
