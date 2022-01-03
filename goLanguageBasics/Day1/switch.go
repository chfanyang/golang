package main

import "fmt"

func main() {
	var yes string
	fmt.Print("是否有卖西瓜的：")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("卖十个包子")
	switch yes {
	case "y", "Y":
		fmt.Println("买一个西瓜")
		//case "Y":
		//	fmt.Println("买一个西瓜")
	}

	var score int
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("不及格")
	}
}
