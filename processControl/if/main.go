package main

import "fmt"

func main() {
	//可以在if条件里赋初值
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
