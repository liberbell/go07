package main

import "fmt"

func main() {
	slice1 := []int{8, 9, 10}
	slice2 := append(slice1, 11, 12)
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice2, slice3)
}
