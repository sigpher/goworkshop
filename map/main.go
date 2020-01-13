package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "how do you do"
	s1Slice := strings.Split(s1, " ")

	countMap := make(map[string]int)

	for _, v := range s1Slice {
		// countMap[v] = countMap[v] + 1
		countMap[v]++
	}

	for k, v := range countMap {
		fmt.Printf("%s\t数量(%d)\n", k, v)
	}
}
