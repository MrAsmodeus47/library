package main

import (
	"fmt"
)

type item string

func assembly(line []item) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)

	}
}

func main() {
	assemblyLine := []item{"bdshadhjj", "dhsavhdgsajhdb", "shahgvduayhjwdh"}
	//assemblyLine := make([]item, 3)
	//assemblyLine[0] = "aaaaaaaaaaa"
	//assemblyLine[1] = "bbbbbbbbbbb"
	//assemblyLine[2] = "ccccccccccc"
	assembly(assemblyLine)
	assemblyLine = append(assemblyLine, "eeeeeeee", "ggggggggggg")
	assembly(assemblyLine)
	assemblyLine = assemblyLine[3:]
	assembly(assemblyLine)
	//fmt.Println(assembly(assemblyLine))

}
