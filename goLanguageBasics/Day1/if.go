package main

import (
	"fmt"
)

func main() {
	fmt.Print("是否有卖西瓜的：")
	var yes string
	fmt.Scan(&yes)
	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")
	if yes == "Y" || yes == "y" {
		fmt.Println("再买一个西瓜")
	}
	fmt.Println("老公的想法：")

	if yes == "Y" || yes == "y" {
		fmt.Println("买一个包子")
	}
	if yes != "Y" && yes != "y" {
		fmt.Println("买十个包子")
	}

	//if else
	var score int
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("不及格")
	}

}
