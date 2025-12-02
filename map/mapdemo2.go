package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//func Printserverstatus(){
//	for _,status := range serverstatus{
//		Println(status)
//	}
func Printserverstatus(servers map[string]int) {
	fmt.Println("\nhere", len(servers), "servers")
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("aaaaa sht here we go again")
		}

	}
	fmt.Println(stats[Online], "server are online")
	fmt.Println(stats[Offline], "server are offline")
	fmt.Println(stats[Maintenance], "serve are under maintenance")
	fmt.Println(stats[Retired], "server are retirde")
}

func main() {
	servers := []string{
		"korea",
		"japone",
		"kowitt",
		"turkey",
		"iran",
		"KSA",
		"china",
	}
	serverstatus := make(map[string]int)
	for _, server := range servers {
		serverstatus[server] = Online
	}
	Printserverstatus(serverstatus)
	serverstatus["iran"] = Retired
	serverstatus["kowitt"] = Offline
	Printserverstatus(serverstatus)
	for server, _ := range serverstatus {
		serverstatus[server] = Maintenance

	}
	Printserverstatus(serverstatus)
}
