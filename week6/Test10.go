package main

import "fmt"

func main() {
	score := 75
	if score > 80 {
		fmt.Println("A")
	} else if score > 70 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("C")
	}
}
