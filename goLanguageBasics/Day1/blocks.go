package main

import (
	"fmt"
)

func main() {
	/*	作用域：定义标识符可以使用的范围
		在GO中用{}来定义作用域范围
		使用原则：子语句块可以使用父语句块中的标识符*/
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		//已存在的变量可以在语句块中重新定义，会重新赋值
		outer = 10
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
	}
}
