package main

import (
	"fmt"
)

func main() {
	shopinglist := map[string]int{
		"salad":     1,
		"soda":      2,
		"olive oil": 1,
		"shipes":    7,
		"khiar":     3,
	}
	shopinglist["khiar"] += 1
	fmt.Println(shopinglist)
	delete(shopinglist, "soda")
	fmt.Println(shopinglist)
	fmt.Println("we need", shopinglist["khiar"], "khiar")
	_, found := shopinglist["soda"]
	if !found {
		fmt.Println("soda deos not found")
	}
	cereal, ok := shopinglist["cereal"]
	fmt.Println("need jdsndajndjsankjdkna?")
	if ok {
		fmt.Println("aaaaaaaaaaaaaaaaa")
	} else {
		fmt.Println("hehhehe", cereal)
	}
	//	tottal := len(shopinglist)
	//	fmt.Println(tottal)
	bigJamshid := 0
	for _, jamshid := range shopinglist {
		bigJamshid += jamshid

	}
	fmt.Println("bigjamshid", bigJamshid)
}
