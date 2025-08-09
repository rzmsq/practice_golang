package main

import (
	"fmt"
)

func main() {
	a, b := [3]int{1, 2, 3}, [3]int{3, 2, 1}
	c := make([]int, 0)

	c = append(c, a[:]...) // Append all elements of a
	c = append(c, b[:]...) // Append all elements of b

	fmt.Println("slice:", c)
	c[0] = 100

	d := [6]int{}
	copy(d[:], a[:])
	copy(d[len(a):], b[:])
	fmt.Println("d:", d)

	e := make([]int, 0)
	e = append(e, a[:]...) // Append all elements of a
	e = append(e, b[:]...) // Append all elements of b

	f := [6]int{}

	copy(f[:], e[:])
	copy(f[len(e):], c[:])
	fmt.Println("f:", f)
}
