package main

import "fmt"

type Rect struct {
	Width  int
	Length int
}

func main() {
	// fmt.Println(Rect{7, 8})
	r := Rect{4, 6}
	p := &r
	p.Width = 8
	fmt.Println(r.Width)
}
