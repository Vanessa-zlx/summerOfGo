package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var e string
	e, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	e = strings.Trim(e, "\n\r")

	eRunes := []rune(e)
	//fmt.Println(len(e), len(eRunes))
	valuence := 1.0
	sideChange := 1.0
	product := 0.0
	factor := 0.0
	var temp []rune
	var x string
	for i := 0; i < len(eRunes); {
		//fmt.Printf("i: %v\n", i) //2x + 5 = 10
		if eRunes[i] == []rune("+")[0] || eRunes[i] == []rune(" ")[0] || unicode.IsLetter(eRunes[i]) {
			i++
		} else if eRunes[i] == []rune("-")[0] {
			valuence = -1.0
			i++
		} else if eRunes[i] == []rune("=")[0] {
			sideChange = -1
			i++
		} else if unicode.IsDigit(eRunes[i]) {
			for i < len(eRunes) && unicode.IsDigit(eRunes[i]) {
				temp = append(temp, eRunes[i])
				i++
			}
			toaddI, _ := strconv.Atoi(string(temp))
			toaddF := float64(toaddI)
			temp = nil
			if i < len(eRunes) && unicode.IsLetter(eRunes[i]) {
				factor += toaddF * valuence * sideChange
				if x == "" {
					x = string(eRunes[i])
				}
			} else {
				product += toaddF * valuence * sideChange
			}
			valuence = 1.0
			//fmt.Println(product, factor)
			if i == len(eRunes) {
				break
			}
		}
	}
	answer := -1.0 * product / factor
	answerI := int(answer*1000000) / 1000
	//fmt.Println(answerI)
	answer = float64(answerI) / 1000
	fmt.Printf("%s=%.3f", x, answer)
}
