package main

import "fmt"

func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	fmt.Println(sum(arr))
	returnTwoIndex(arr)

}

// sum 对数组内的整数进行相加并返回
func sum(arr [5]int) (total int) {
	for _, value := range arr {
		total = total + value
	}
	return
}

// returnTwoIndex if index1 + index2 = 8,return two index
func returnTwoIndex(arr [5]int) {
	for k1, v1 := range arr {
		for k2, v2 := range arr {
			if v1+v2 == 8 {
				if k1 < k2 {
					fmt.Println(k1, k2)
				}
			}
		}
	}
}
