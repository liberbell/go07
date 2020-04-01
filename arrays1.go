package main

import "fmt"

func main() {
	var z [3]string
	z[0] = "Hello"
	z[1] = "My"
	z[2] = "World"

	fmt.Println(z)

	var y [4]int
	y[3] = 10
	fmt.Println(y)
}
