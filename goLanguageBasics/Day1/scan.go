package main

import "fmt"

func main() {
	var name string
	fmt.Printf("请输入姓名:")
	//将输入的内容赋值给对应的变量
	fmt.Scan(&name)
	fmt.Println("姓名:" + name)
}
