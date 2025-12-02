package main

import "fmt"

type Counter struct {
	hite int
}

func inceremnt(counter *Counter) {
	counter.hite += 1
	fmt.Println("Counter", counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new
	inceremnt(counter)
}
func main() {
	counter := Counter{}

	hello := "rrrrrrrrrr"
	world := "xdxdxxxdxxxxd"
	fmt.Println(hello, world)
	replace(&hello, "fufuf", &counter)
	fmt.Println(hello, world)
	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "hdfuysygu", &counter)
	fmt.Println(phrase)
}
