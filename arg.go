package main

import "fmt"

func family(fname string, age int) {
	fmt.Printf(" %v \t %v \n", fname, age)
}
func familyc(a int, b int, c int) (result int, sa int) {
	result = a + b + c
	sa = a + b
	return
}
func main() {
	family("guru", 22)
	family("prasanth", 49)
	family("ammu", 50)
	fmt.Println(familyc(10, 22, 50))
}
