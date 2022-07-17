package main

import "fmt"

func main() {
	m := 5
	switch m {
	case 1, 3, 5, 7, 9:
		//case 后面可以有多个值或者写 表达式
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}
}
