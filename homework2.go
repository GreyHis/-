package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}
func Calculator(a int, b int, CMD func(int, int) int) int {
	return CMD(a, b)
}
func main() {
	var c string
	var a, b int
	fmt.Println("请输入两个数和运算符号")
	fmt.Scanf("%d %d %s", &a, &b, &c)
	switch c {
	case "+":
		fmt.Println(Calculator(a, b, sum))
	case "-":
		fmt.Println(Calculator(a, b, sub))
	case "*":
		fmt.Println(Calculator(a, b, multiply))
	case "/":
		fmt.Println(Calculator(a, b, divide))
	}

}
