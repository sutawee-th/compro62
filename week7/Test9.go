package main

func makeEven() func() int {
	even := 0
	return func() int {
		even = even + 2
	}
}
