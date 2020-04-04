package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	fmt.Println(x)
}
