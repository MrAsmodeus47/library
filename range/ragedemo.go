package main

import "fmt"

func main() {
	slice := []string{"book of the dead", "laier laier", "witcher"}
	for i, slcvalue := range slice {
		fmt.Println(i, slcvalue, ":")
		for _, strch := range slcvalue {
			fmt.Printf("%q\n", strch)

		}
	}
}
