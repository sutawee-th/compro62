package main

import "fmt"

func handlePanic() {
	text := recover()
	fmt.Println(text)
}

func main() {

}
