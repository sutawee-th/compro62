package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("Hello")
	b.WriteString(" ")
	b.WriteString("Go")
	fmt.Println(b.String())
}
