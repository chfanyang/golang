package main

import "fmt"

func main() {
	//breack
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue //跳出本次循环，进入下次循环
		}
		fmt.Println(i)
	}
	//continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			break //结束循环
		}
		fmt.Println(i)
	}
}
