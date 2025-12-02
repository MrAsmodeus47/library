package main

import "fmt"

type Space struct {
	occupied bool
}
type Parkinglote struct {
	spaces []Space
}

func occupiedspace(lot *Parkinglote, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true

}
func (lot *Parkinglote) occupiedspace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}
func (lot *Parkinglote) vacadeSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false ///////////////

}

func main() {
	lot := Parkinglote{spaces: make([]Space, 10)}
	fmt.Println("initial:", lot)
	lot.occupiedspace(1)   //////////this two are do same thing
	occupiedspace(&lot, 2) ////////this two are do same thing
	fmt.Println("after so some sht:", lot)
	lot.vacadeSpace(2)
	fmt.Println("after vaction :", lot)

}
