package main

import "fmt"

const (
	Active   = true
	Diactive = false
)

type scurityTag bool
type Item struct {
	name string
	tag  scurityTag
}

func activate(tag *scurityTag) {
	*tag = Active
}
func diactivated(tag *scurityTag) {
	*tag = Diactive
}
func ckeckout(item []Item) {
	fmt.Println("ckenking out...")
	for i := 0; i < len(item); i++ {
		diactivated(&item[i].tag)

	}
}

func main() {
	shirt := Item{"Shirt", Active}
	dress := Item{"Dress", Active}
	jacket := Item{"Jachet", Active}
	hoodi := Item{"Hoodi", Active}
	items := []Item{shirt, dress, jacket, hoodi}
	fmt.Println(items)
	diactivated(&items[2].tag)
	fmt.Println("item 3 is diactivated", items)
	ckeckout(items)
	fmt.Println("ckecket out", items)

}
