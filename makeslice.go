package main

func main() {
	a := make([]int, 4)
	printSl(("a", a))
}

func printSl(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}