package main

import "fmt"

func main() {
	x := make(map[string]int)
	fmt.Println(len(x))
	x["key"] = 99
	fmt.Println(len(x))
	fmt.Println(x)
}
