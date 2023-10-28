package main

import "fmt"

func sushu(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}

	}
	return true

}
func main() {
	var i int
	fmt.Print("请输入一个整数：")
	fmt.Scanf("%d", &i)
	result := sushu(i)
	if result == true {
		fmt.Printf("整数%d是素数", i)

	} else {
		fmt.Printf("整数%d不是素数", i)
	}
}
