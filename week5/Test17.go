package main

import "fmt"

func main() {
	humans := []string{"Bulma", "chi-chi", "Videl"}
	names := []string{"Goku", "Gohan"}
	names = append(names, humans...)
	fmt.Println(names)
}
