package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reder := bufio.NewReader(os.Stdin)
	name, _ := reder.ReadString('\n')
	//second method
	var age int
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Printf("error scanning age")
	}
	fmt.Printf("age: %d\n", age)
	fmt.Print("name:", name)
}
