package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i = i + 1
		if i >= 3 {
			break
		}
	}
}
