package main

import (
	"fmt"
)

const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

const (
	Admin      = 50
	Manager    = 40
	Contractor = 30
	Member     = 20
	Guest      = 10
)

func weekday(today int) bool {
	return today <= 4
}

func accessGranted() {
	fmt.Println("welcome")
}

func accessDenied() {
	fmt.Println("fck off")
}

func main() {
	var today, role int
	fmt.Print("what day is to day?")
	fmt.Scanln(&today)

	fmt.Print("your rank")
	fmt.Scanln(&role)

	if role == Admin || role == Manager {

		accessGranted()
	} else if role == Contractor && !weekday(today) { //Saturday || Sunday
		accessGranted()
	} else if role == Member && weekday(today) { //Monday || Tuesday || Wednesday|| Thursday|| Friday
		accessGranted()
	} else if role == Guest && (today == Wednesday || today == Friday || today == Monday) { //weekday(today)==2 || weekday(today)==4 || weekday(today)==0
		accessGranted()
	} else {
		accessDenied()
	}
}
