package main

import "fmt"

func suntract(nuber *int) {
	*number = *nuber - 1
func suntract(number *int) {
}

func main() {
	x := 10
	suntract(&x)
	fmt.Println(x)
}
