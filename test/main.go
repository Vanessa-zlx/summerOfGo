package main

import (
	"fmt"
	//"unicode"
)

func main() {
	var x string
	if x == "" {
		fmt.Println("**")
	}
	// a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// a = strings.Trim(a, "\n\r")
	// strings.Contains(a, "\r")
	// fmt.Printf("len(a): %v\n", len(a))
	// fmt.Println(strings.LastIndex(a, " "))
	// fmt.Printf("len(a): %v\n", len(a))
	// var a []rune
	// b := []rune("15246")
	// for i := 0; i < len(b); i++ {
	// 	if unicode.IsDigit(b[i]) {
	// 		a = append(a, b[i])
	// 	}
	// }
	// x, _ := strconv.Atoi(string(a))
	// fmt.Println(a, x+1)
	//a := "fff1你好5哇"
	//fmt.Printf("%T\n", []rune("a")[0])

	//b := []rune(a)
	// for i := 0; i < len(b); i++ {
	// 	fmt.Println(string(b[i]))
	// }
	// for _, v := range a {
	// 	fmt.Println(unicode.IsDigit(v))
	// }
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%T\n", a[i])
	// }
	// b := []rune(a)
	// for i := 0; i < len(b); i++ {
	// 	fmt.Println(b[i])
	// }

	//fmt.Printf("%T,%T", a[0], b[0])
	//fmt.Println(unicode.IsDigit(rune(a[0])))
	//fmt.Println(b)
}
