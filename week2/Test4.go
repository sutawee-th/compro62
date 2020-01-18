package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("Sutawee")
	b.WriteString(" ")
	b.WriteString("Taenkratok")
	fmt.Println(b.String())
}
