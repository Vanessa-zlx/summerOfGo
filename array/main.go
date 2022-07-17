package main

import "fmt"

func main() {
	//1.初始化
	var a [3]int
	var b [5]int
	//1.1长度必须是常量
	fmt.Printf("%T,%T", a, b)
	//1.2不同长度的数组类型不同
	fmt.Println(a, b)
	//2.赋值
	//2.1用大括号赋值可以省略长度
	var c = [...]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(c)
	//2.2索引赋值
	var d = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	//3遍历
	//3.1for i
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
	//3.2for range
	for index := range d {
		fmt.Println(index)
	}
	for index, value := range d {
		fmt.Println(index, value)
	}
	for _, v := range d {
		fmt.Println(v)
	}

	//4二维数组
	//4.1赋值
	city := [...][2]string{ //只有外层长度可以省略
		{"北京", "上海"},
		{"西安", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(city)
	//4.2遍历
	for _, v := range city {
		fmt.Println(v)
	}
	for _, v1 := range city {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//5值传递
	values := [3]int{1, 2, 3}
	values = f1(values)
	fmt.Println(values)
}

func f1(a [3]int) [3]int {
	a[0] = 100
	return a
}
