package main

import "fmt"

func prinSlice(title string, slice []string) {
	fmt.Println("")
	fmt.Println("-----", title, "------,")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)

	}

}
func main() {
	newLoc := ""
	fmt.Println("where the fck are you going ?")
	fmt.Scanln(&newLoc)
	route := []string{"office", "park", "home", "shop", "salon"}
	prinSlice("road 1", route)
	route = append(route, newLoc)
	prinSlice("road 2", route)
	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited", ";")
	fmt.Println("remaining loction", route[2:])

}
