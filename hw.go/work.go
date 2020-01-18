package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i <= 3000; i++ {
		file, err := os.Create(fmt.Sprintf("%v.txt", i))

		if err != nil {
			return
		}
		defer file.Close()

		for s := 0; s < 1000; s++ {
			file.WriteString("Sutawee\n")
		}
		for s := 0; s < 1000; s++ {
			file.WriteString("Taenkratok\n")
		}
		for s := 0; s < 1000; s++ {
			file.WriteString("39\n")
		}
	}
	fmt.Println(time.Since(start))
}
