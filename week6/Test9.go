package main

import "fmt"

func main() {
	for i := 1; i <= 10; i = i + 1 {
		if i%2 == 0 {
			fmt.Println(i, "even")
		}
	}
}
