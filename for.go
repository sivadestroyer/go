package main

import (
	"fmt"
)

func main() {
	k := [10]int{1, 2, 3, 4, 5}
	for x := 1; x < 6; x++ { // for iteration
		fmt.Println(k[x])
	}
	for idx, val := range k { //for in different ways
		fmt.Println("\n", idx, val)
	}
}
