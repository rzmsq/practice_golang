package main

import (
	"fmt"
	"os"
)

type Arg struct {
	Index int
	Value string
}

func main() {
	argsSlice := make([]Arg, 0, len(os.Args)-1)
	for i := 1; i < len(os.Args); i++ {
		argsSlice = append(argsSlice, Arg{i, os.Args[i]})
	}

	fmt.Println(argsSlice)
}
