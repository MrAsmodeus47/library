package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		dividedBy3 := i%3 == 0
		dividedBy5 := i%5 == 0
		if dividedBy3 && dividedBy5 {
			fmt.Println("fizzbuzz")

		} else if dividedBy3 {
			fmt.Println("fizz")
		} else if dividedBy5 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
