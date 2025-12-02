package main

import "fmt"

type Items struct {
	Name  string
	Price int
}

func cashier(item Items) {}

func main() {
	//shopingList[0].Items=9
	//shopingList[1].Items=12
	//shopingList[2].Items=10
	//shopingList[3].Items=11

	shopingList := [4]Items{
		{Price: 9, Name: "khiar"},
		//{Price: 9},
		{Price: 12, Name: "mmmmmoooooooossss"},
		//{Price: 12},
		{Price: 10, Name: "mast"},
		//{Price: 10},
		//{Price : 11,Name: "noshaba"},
		//{Price: 11},
	}
	//shopingList[0].Price = 9
	//shopingList[1].Price = 12
	//shopingList[2].Price = 10
	//shopingList[3].Price = 11
	fmt.Println(shopingList)

}
