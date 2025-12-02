package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int
type Liftpicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("motorcycle : %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car : %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck : %v", string(t))
}
func (m Motorcycle) PickLift() Lift {
	return SmallLift
}
func (c Car) PickLift() Lift {
	return StandardLift
}
func (t Truck) PickLift() Lift {
	return LargeLift
}
func sendToLift(lp Liftpicker) {
	switch lp.PickLift() {
	case SmallLift:
		fmt.Printf("%v need a small left\n", lp)
	case StandardLift:
		fmt.Printf("%v need a standard left\n", lp)
	case LargeLift:
		fmt.Printf("%v need a large left\n", lp)
	}
}
func main() {
	car := Car("golfmini")
	truck := Truck("T-rex")
	motorcycle := Motorcycle("rs200")
	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)

}
