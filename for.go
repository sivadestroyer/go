package main

import (
	"fmt"
)

func main() {
	k := [100]int{1, 2, 3, 4, 5}
	for x := 1; x < 6; x++ {
		fmt.Println(k[x])
	}
}
