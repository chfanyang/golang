package main

import "fmt"

func main() {
	var sum int
	//语法一
	//初始化子语句; 条件子语句; 后置子语句
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//语法二
	i, sum := 0, 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
	//语法三（死循环）
	for {
		i++
		fmt.Println(i)
	}
	//for range
	desc := "我爱中国"
	for i, ch := range desc {
		fmt.Printf("%d %T %q\n", i, ch, ch)
	}
}
