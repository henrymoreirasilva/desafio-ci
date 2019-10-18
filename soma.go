package main

import (
	"fmt"
)

func main() {
	var r int
	
	r = Soma(5, 5)

	fmt.Println(r)
}

func Soma(a int, b int) int {
	return a + b
}