package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
		// fmt.Println(a)
	}
	fmt.Println(a)

	arr := [...]int{3, 7, 8, 9, 1}
	reverseArr := arr[:]
	sort.Ints(reverseArr)
	fmt.Println(reverseArr)

}

func intSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
