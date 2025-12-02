package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}
func main() {
	var dice, sides int
	rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("how many dice do you want")
	fmt.Scanln(&dice)
	fmt.Println("how many sides your dice have")
	fmt.Scanln(&sides)
	rolls := 1
	for R := 1; R <= rolls; R++ {
		sum := 0
		for D := 1; D <= dice; D++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println(R, "Dice#", D, ":", rolled)
		}
		fmt.Println("Totall Rolled", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes")
		case sum == 7:
			fmt.Println("Lucky")
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}

	}

}
