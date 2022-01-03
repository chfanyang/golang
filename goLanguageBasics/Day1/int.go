package main

import "fmt"

func main() {
	var age int = 21
	//字面量：十进制，八进制，十六进制
	fmt.Printf("%T, %d\n", age, age)
	fmt.Println(0x777)
	fmt.Println(0777)
	//操作
	//算数运算(+, -, *, %)
	fmt.Println(1 + 2)
	fmt.Println(2 - 2)
	fmt.Println(2 * 2)
	fmt.Println(9 % 2)
	//自增(++)，自减(--)
	a := 1
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	//关系运算(<,<=,>,>=,==,!=)
	fmt.Println(2 > 3)
	fmt.Println(2 < 3)
	fmt.Println(2 <= 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 != 3)
	fmt.Println(2 == 3)
	//位运算
	//赋值运算符(= ,+= ,-= ,*= ,/= ,%=  ,<<=  ,>>= ,&= ,^= ,|= )
	a = 1
	a += 1
	fmt.Println(a)

	//类型转换
	//在go中没有自动类型转换，只有强制类型转换
	var intA int = 10
	var intB uint = 3
	fmt.Println(intA + int(intB))
	fmt.Println(uint(intA) + intB)
	//大 -> 小 转换可能出现溢出
	var intC int = 0xFFFF
	fmt.Println(intC, uint8(intC), int8(intC))
}
