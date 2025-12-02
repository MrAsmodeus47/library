package main

import "fmt"

type Room struct {
	Name    string
	Cleaned bool
}

func roomCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.Cleaned {
			fmt.Println(room.Name, "is clean")
		} else {
			fmt.Println(room.Name, "need to be clean")
		}
	}
}

func main() {
	//var roomNum int
	//fmt.Println("enter the room number")
	//fmt.Println(&roomNum)
	rooms := [...]Room{
		{Name: "office-1"},
		{Name: "office-2"},
		{Name: "office-3"},
		{Name: "office-4"},
	}
	roomCleanliness(rooms)
	fmt.Println("cleanung in progress")
	rooms[1].Cleaned = true
	rooms[3].Cleaned = true
	roomCleanliness(rooms)
}
