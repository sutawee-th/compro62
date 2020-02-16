package main


type myError struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func say(word string) error {
	if word == "hi" {
		return myError{"can't say hi"}
	}
	return nil
}
func main() {
	e1 := say("hello")
	fmt.Println(e1)
	e2 := say("hi")
	fmt.Println(e2)
}
