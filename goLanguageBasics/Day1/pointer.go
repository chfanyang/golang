package main

import "fmt"

func main() {
	a := 1
	b := a
	b = 3
	fmt.Println(a, b)
	//指针
	//将内存地址赋值给指针类型的变量
	var cc *int = &a
	c := &a
	fmt.Printf("%T %T\n", c, cc)
	fmt.Println(*cc)
	*cc = 2
	fmt.Println(a)

}
