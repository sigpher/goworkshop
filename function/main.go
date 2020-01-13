package main

import (
	"fmt"
	"strings"
)

// Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。

func main() {
	namestring := "aaaEUUU,Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth,choi,uuu"
	names := strings.Split(namestring, ",")
	distribution := make(map[string]int, len(names))
	coin := 50

	for _, username := range names {
		nameSlice := []byte(strings.ToLower(username))
		userCoin := 0
		for _, letter := range nameSlice {
			switch string(letter) {
			case "e":
				if coin >= 1 {
					userCoin = userCoin + 1
					coin = coin - 1
				}
			case "i":
				if coin >= 2 {
					userCoin = userCoin + 2
					coin = coin - 2
				}
			case "o":
				if coin >= 3 {
					userCoin = userCoin + 3
					coin = coin - 3
				}
			case "u":
				if coin >= 4 {
					userCoin = userCoin + 4
					coin = coin - 4
				}

			}
			distribution[username] = userCoin
		}
	}
	fmt.Println(distribution)
	fmt.Println(coin)

}
