package main

import (
	"fmt"
	"os"
	"strconv"
)

type Arr2x2 [2][2]int

func Add(a, b Arr2x2) Arr2x2 {
	c := Arr2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func (a *Arr2x2) Add(b Arr2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] += b[i][j]
		}
	}
}

func (a *Arr2x2) Subtract(b Arr2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] -= b[i][j]
		}
	}
}

func (a *Arr2x2) Multiply(b Arr2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		os.Exit(1)
	}

	k := [8]int{}
	for index, i := range os.Args[1:] {
		var err error
		k[index], err = strconv.Atoi(i)
		if err != nil {
			os.Exit(1)
		}
	}
	a := Arr2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := Arr2x2{{k[4], k[5]}, {k[6], k[7]}}

	fmt.Println("Traditional Add:", Add(a, b))
}
