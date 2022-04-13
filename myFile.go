package main

import (
	"fmt"
)

func main() {
	//var div int
	var b = 99
	var a = 3
	//div = b / b
	division(b, a)
	b = 50
	a = 5
	division(b, a)
	division(b, a)
}
func division(firstNumber int, secondNumber int) {
	var div int
	div = firstNumber / secondNumber
	fmt.Println(div)
}
