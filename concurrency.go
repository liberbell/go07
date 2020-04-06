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

	go msg()
	fmt.Println("\nMessage from func main, I`m finished.")
	time.Sleep(time.Millisecond * 7500)
}

func msg() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i > 3 {
			fmt.Println(i, "seconds... yawn")
		} else {
			fmt.Println(i, "seconds")
		}
	}
	fmt.Println("\nMessage from func msg: I`m finished.")
}
