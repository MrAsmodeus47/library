package main

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func area(rec Rectangle) int {
	return rec.Length * rec.Width

}
func perimeter(rec Rectangle) int {
	return (rec.Length + rec.Width) * 2

}
func main() {
	rect := Rectangle{}
	fmt.Println("enter the lengh")
	fmt.Scanln(&rect.Length)
	fmt.Println("enter the widgh")
	fmt.Scanln(&rect.Width)~
	fmt.Println("area :", area(rect), "perimeter :", perimeter(rect))
}
