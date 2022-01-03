package main

import "fmt"

func main() {
	//float32，float64
	//定义float64类型变量
	var height float64 = 1.68
	fmt.Printf("%T, %f\n", height, height)
	//字面量 ：十进制碧松暗示法 科学记数法表示法
	var weight float64 = 13.05e1
	fmt.Println(weight)

	//操作(float类型在存储时，会存在精度损耗)
	//算术运算(+, -, *, /, ++, --)
	fmt.Println(1.1 - 1.2)
	fmt.Println(1.1 + 1.2)
	fmt.Println(1.1 * 1.2)
	fmt.Println(1.1 / 1.2)
	height++
	fmt.Println(height)
	height--
	fmt.Println(height)
	//关系运算(> >= < <=)
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 >= 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)
	//赋值运算(=, +=, -=, /=, *=)
	height += 0.5
	fmt.Println(height)

	fmt.Printf("%T, %T\n", 1.1, float32(1.1))

	fmt.Printf("%5.2f\n", height)
}
