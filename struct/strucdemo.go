package main

import "fmt"

type Passenger struct {
	Name         string
	Ticketnumber int
	Boarded      bool
}
type Bus struct {
	Vip Passenger
}

func main() {
	mmad := Passenger{"mmadof", 32124543, false}
	fmt.Println(mmad)
	var (
		mahshid = Passenger{"mahshid", 3213231, true}
		asqar   = Passenger{Name: "asqar", Ticketnumber: 555555}
	)
	fmt.Println(asqar, mahshid)
	var reza Passenger
	reza.Name = "frank"
	reza.Ticketnumber = 8748782356
	fmt.Println(reza)

	mmad.Boarded = true
	asqar.Boarded = true
	if asqar.Boarded {
		fmt.Println("he has come")

	}
	if asqar.Boarded {
		fmt.Println(asqar.Name, "has boarded")
	}
	reza.Boarded = true
	bus := Bus{reza}
	fmt.Println(bus)
	fmt.Println(bus.Vip.Name, "is here")

}
