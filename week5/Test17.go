package main

func main() {
	humans := []string{"Bulma", "chi-chi", "Videl"}
	names := []string{"Goku", "Gohan"}
	names = append(names, humans...)
}
