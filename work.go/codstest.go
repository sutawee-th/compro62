package main

import (
	"log"
	"os"
)

func readCurrentDir() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("failed opening directory")
	}
	defer file.Close()
}

func main() {
	readCurrentDir()
}
