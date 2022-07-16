package main

import "fmt"

var x = 100
var y string = "小王子"

//函数外部变量声明赋值必须有关键字

func main() {
	fmt.Println(x, y)
	//1注释
	//1.1单行注释
	fmt.Println("hello world !")

	//2声明
	//2.1标准声明
	var name string
	var age int
	var isOK bool
	fmt.Println(name, age, isOK)
	//2.2批量声明
	var (
		a string  = "a"
		b int     = 2
		c bool    = false
		d float32 = 1.1
	)
	fmt.Println(a, b, c, d)
	//2.3类型推导 --- 可以用在全局
	var name1 = "zhulx"
	fmt.Println(name1)
	//2.4短变量声明  --- 只能用在函数内部
	m := 10
	fmt.Println(m)
	/*2.5匿名变量  --- 哑元变量
	  用  _  表示，不分牌内存*/

	//3常量
	const pi = 3.14159
	const e = 2.71828
	fmt.Println(pi, e)
	//3.1支持批量声明
	//3.2相同值常量赋值简写
	const (
		n1 = 10
		n2
		n3
	)
	fmt.Println(n1, n2, n3)
	//3.3计数器iota
	//可以简写
	//遇到const变0
	const (
		m1 = iota
		//换行iota值+1
		m2 = iota
		_
		//匿名变量会跳过一次
		m4 = iota
	)
	fmt.Println(m1, m2, m4)
	const (
		k1 = iota
		k2 = 200
		//可以插队，但是iota仍然加一
		_
		k4 = iota
		k5 = 1 << (10 * (k4 - 2))
		//可以参与运算
	)
	fmt.Println(k1, k2, k4, k5)
}
