package main

import "fmt"

func main() {
	//布尔类型，表示真假
	//标识符bool
	//可选值true.false
	var zero bool
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
	//bool操作
	//逻辑运算（与&&, 或||， 非!）
	//与运算
	fmt.Println("&&")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != isGirl)
}
