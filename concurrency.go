package main

import (
	"fmt"
	"time"
)

func f(msg int) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	// go f("value of i")
	// var input string
	// fmt.Scanln(&input)

	msg()
	fmt.Println("\nMessage from func main, I`m finished.")
	time.Sleep(time.Millisecond * 2500)
}
