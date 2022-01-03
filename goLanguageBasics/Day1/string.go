package main

import "fmt"

func main() {
	//"" -> 可解释的字符串
	//`` ->原生字符串
	//特殊字符 \n \r \n \t \f \b \v
	var name string = "a\tb"
	var desc string = `我来\t自中国`
	fmt.Println(name)
	fmt.Println(desc)
	//操作
	//算术运算符: + (连接)
	fmt.Println("我叫" + "小明")
	//关系运算 (== != > >= < <=)
	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" < "bb")
	fmt.Println("ab" <= "bb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" >= "bb")
	//赋值
	s := "我叫"
	s += "小明"
	fmt.Println(s)
	//字符串定义的内容必须是ASCII码
	//索引 0 - n-1 (n 代表字符串的长度)
	desc = "abcdef"
	fmt.Printf("T %c\n", desc[0], desc[0])
	//切片[start:end] start end-1
	fmt.Printf("%T %s\n", desc[0:2], desc[0:2])
	//获取字符串长度
	fmt.Println(len(desc))
}
