package main

import "fmt"

func main() {
	//1.算数运算符
	fmt.Println(5 / 2)
	a := 10
	//++,--不是运算符而是语句，只能单独写
	a++
	//赋值运算符  a  xx= b
	a += 5
	fmt.Println(a)
	//关系运算符 ==
	//逻辑运算符 && || ！
	//位运算符   & | ^ << >>
	fmt.Println(1&5, 1|5, 1^5) //001,101
	fmt.Println(1<<2, 4>>2)
}
