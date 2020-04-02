package main

import "fmt"

func main() {
	a := make([]int, 4)
	printSl("a", a)
	b := make([]int, 0, 4)
	printSl("b", b)
	c := b[:1]
	printSl("c", c)
	d := c[2:4]
	printSl("d", d)
}

func printSl(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
