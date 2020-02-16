package main

type myError struct {
	error string
}

func say(word string) error {
	if word == "hi" {
		return myError{"can't say hi"}
	}
	return nil
}
func main() {
