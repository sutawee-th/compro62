package main

func say(txt string) {
	for i = 0; < 3; i++ {
		fmt.Println(i, " : ", txt)
	}
}

func main() {
	go say("Hello")
	go say("Hi")
}
