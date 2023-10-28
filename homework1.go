package main

import (
	"fmt"
	"math"
)

func area(r float64) float64 {
	area := r * r * math.Pi
	return area
}
func main() {
	fmt.Print("请输入半径：")
	var r float64
	fmt.Scanf("%f", &r)
	area := area(r)
	fmt.Printf("面积为:%.2f", area)

}
