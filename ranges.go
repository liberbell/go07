package main

import "fmt"

func main() {
	for n, p := range pow {
		fmt.Printf("2**%d = %d\n", n, p)
	}
}
