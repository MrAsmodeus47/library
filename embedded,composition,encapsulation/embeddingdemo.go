package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return "woooof" + d.Name }
func (c Cat) Speak() string { return "meeeowwwww" + c.Name }

type Farm struct {
	Animals []Animal
}

func (f *Farm) speakall() []string {
	out := make([]string, 0, len(f.Animals))
	for _, a := range f.Animals {
		out = append(out, a.Speak())
	}

	return out
}

func main() {

	f := &Farm{
		Animals: []Animal{Dog{Name: "rex"}, Cat{Name: "losy"}},
	}
	fmt.Println(f.speakall())

}
