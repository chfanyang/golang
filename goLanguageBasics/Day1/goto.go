package main

import "fmt"

func main() {
	var yes string
	fmt.Print("有卖西瓜的吗：")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")

	if yes != "y" && yes != "Y" {
		fmt.Println(yes)
		goto END
	}
	fmt.Println("买一个西瓜")

END:
}
