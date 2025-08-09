package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f: ", &f)

	fP := &f
	fmt.Println("Memory address of fP: ", fP)
	fmt.Println("Value of fP: ", *fP)

	processPointer(fP)
	fmt.Printf("Value of fP: %.2d\n", fP)

	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	var k *aStructure
	fmt.Println(k)
	if k == nil {
		k = new(aStructure)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil")
	}

}
