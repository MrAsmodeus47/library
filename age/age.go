package main

import "fmt"

func main() {
	var age int
	fmt.Println("enter your age")
	fmt.Scanln(&age)
	switch {
	case age == 0:
		fmt.Println("coming soon")
	case 1 >= age && age <= 3:
		fmt.Println("in development")

	case 4 >= age && age <= 12:
		fmt.Println("tefl")
	case 13 >= age && age <= 17:
		fmt.Println("shebhe tefl")
	default:
		fmt.Println("old as fck")

	}
}
