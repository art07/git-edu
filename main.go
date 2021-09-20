package main

import "fmt"

func main() {
	fmt.Println("Test git app.")
	fmt.Println("New msg.")

	fmt.Println("123")
	fmt.Println(myFunc(5, 7))
}

func myFunc(i1, i2 int32) int32 {
	return i1 + i2
}
