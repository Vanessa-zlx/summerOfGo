package main

import "fmt"

func main() {
	//1.基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//2.省略写法,写在外面和里面
	var i = 10
	for i > 0 {
		fmt.Println(i)
		i--
	}
	//3.无限循环
	// for {
	// 	fmt.Println("hello")
	// }
}
