package main

import "fmt"

func main() {
	w := make(map[string]int)
	w["Answer"] = 10
	fmt.Println("The valeu: ", w["Answer"])

	w["Answer"] = 20
	fmt.Println("The value: ", w["Answer"])

	delete(w, "Answer")
	fmt.Println("The value: ", w["Answer"])

	v, ok = w["Answer"]
	fmt.Println("The value: ", v, "Presetn?", ok)
}
