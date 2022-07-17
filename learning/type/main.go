package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//1.整数类型
	/*
		uint8：byte,uint16:short,uint32,uint64:long,
		int8,int16,int32,int64
		unint,int：与系统有关32/64；uintptr：存放指针
	*/
	//2浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	//3布尔值 --- 不能参与运算，非int
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	//4字符串
	//4.1 utf-8编码
	s1 := "hello"
	s2 := "你好北京"
	fmt.Println(s1)
	fmt.Println(s2)
	//4.2转义字符
	fmt.Println("D:\\git\\summerOfGo")
	fmt.Printf("\t制表\n换行aa\n")
	//4.3多行字符串
	s3 := `开始
多行字符串\n\t
\n\t
	4`
	fmt.Println(s3)
	//4.4字符串方法
	s4 := "你好world"
	s5 := "hello世界"
	//长度
	fmt.Println(len(s4), len(s5))
	//拼接
	fmt.Println(s4 + s5)
	s6 := fmt.Sprintf("%s-%s", s4, s5)
	//Sprint:输出到字符串；Fprint：输出到stream；Printf：标准输出
	fmt.Println(s6)
	/*
		strings.Contains(s1,s2);
		strings.Split(s1,s2);
		strings.HasPrefix();string.HasSuffix(s1,s2)
		strings.LastIndex(s1,s2),fmt.Index(s1,s2)
		strings.join(s1,s2)
	*/
	fmt.Println(strings.Contains(s4, "or"))
	s7 := "how do you do ?"
	fmt.Println(strings.Split(s7, " "), strings.Contains(s7, "do"), strings.HasPrefix(s7, "ho"), strings.HasSuffix(s7, " ?"))
	fmt.Printf("%T\n", strings.Split(s7, " "))
	fmt.Println(strings.Index(s7, "do"), strings.LastIndex(s7, "do"))
	s8 := strings.Split(s7, " ")
	fmt.Println(strings.Join(s8, "$"))

	//4.4字符
	//byte uint8 ASCII码
	//rune int32
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("9%T,%T\n", c1, c2)
	s9 := "hello世界"
	for i := 0; i < len(s9); i++ {
		fmt.Printf("%c\n", s9[i])
	}
	for _, v := range s9 {
		fmt.Printf("%c\n", v)
	}
}
