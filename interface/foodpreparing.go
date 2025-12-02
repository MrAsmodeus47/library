package main

import "fmt"

type Preparer interface {
	Preparedish()
}

type Chichen string
type Salad string

func (c Chichen) Preparedish() {
	fmt.Println("cook ckicken")
}
func (s Salad) Preparedish() {
	fmt.Println("prepare salad")
	fmt.Println("add dressing")

}
func prepareDishes(dishes []Preparer) {
	fmt.Println("prepare dishes :")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("---Dish %v---\n", dish)
		dish.Preparedish()
	}
	fmt.Println()
}
func main() {
	dishes := []Preparer{Chichen("Chichen wing"), Salad("sezar salad")}
	prepareDishes(dishes)
}
