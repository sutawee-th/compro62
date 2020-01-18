package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Hello Word", "hi"))
	fmt.Println(strings.ContainsAny("Hello Word", "Hi"))
}
