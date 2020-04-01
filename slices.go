package main

import "fmt"

func main() {
	slice := []int{7, 9, 10, 12}
	fmt.Println("Slice = ", slice)

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("Slice[%d] = %d\n", i, slice[i])
	// }
	fmt.Println("Slice[1:4] = ", slice[1:4])
	fmt.Println("Slice[:3]  ==", slice[:3])
}
