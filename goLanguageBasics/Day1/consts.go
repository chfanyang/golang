package main

import "fmt"

func main() {
	//声明常量
	const NAME string = "kk"
	fmt.Println(NAME)
	//声明常量（省略类型）
	const C1 = 2
	//声明多个类型相同的常量
	const C2, C3 int = 2, 3
	//声明多个类型不同的变量
	const (
		C4 int32   = 1
		C5 float64 = 1.8
	)
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)
	fmt.Println(C4)
	fmt.Println(C5)
	//批量声明相同值常量
	const (
		C6 int = 9
		C7
		C8
	)
	fmt.Println(C6)
	fmt.Println(C7)
	fmt.Println(C8)
	//枚举
	const (
		E1 int = iota
		E2
		E3 = 4
		E4
	)
	fmt.Println(E1, E2, E3, E4)
	//右侧可以是表达式
	const (
		E5 int = iota + 10
		E6
		E7
	)
	fmt.Println(E5, E6, E7)
}
