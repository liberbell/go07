package main

import "fmt"

func main() {
	a := make([]int, 4)
	printSl("a", a)
	b := make([]int, 0, 4)
	printSl("b", b)
}

func printSl(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
