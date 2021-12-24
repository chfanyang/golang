package main

import (
	"fmt"
)

//定义全局变量并赋值
var version string = "cdc"

func main() {
	//定义变量
	/*
		变量名称需要满足标识符命名规则
		1. 必须由非空的Unicode字符串组成、数字、_
		2. 不能以数字开头
		3. 不能是go的关键字
		4. 避免和go预定义的标识符冲突
		5. 建议使用驼峰命名法
		6. 标识符区分大小写
	*/
	var me string
	me = "小明"
	fmt.Println(me)
	//全局变量可以定义后不使用
	//fmt.Println(version)

	//定义多个变量
	var name, user string = "1", "2"
	fmt.Println(name, user)
	//定义多个类型不同的变量
	var (
		age    int
		height float64
	)
	fmt.Println(age, height)
	//声明多个变量同时赋值
	var (
		age1    int     = 1
		height1 float64 = 1.77
	)
	fmt.Println(age1, height1)
	//自动推导变量类型
	var a = 1
	var b = 2
	//自动推导多个变量类型
	var c, d = 3, 4
	fmt.Println(a, b)
	fmt.Println(c, d)
	//简短声明,只能在函数内部使用
	isBoy := true
	fmt.Println(isBoy)
	isBoy, c, d = false, 5, 6
	fmt.Println(isBoy, c, d)
	//交换变量的值
	c, d = d, c
	fmt.Println(c, d)

}
