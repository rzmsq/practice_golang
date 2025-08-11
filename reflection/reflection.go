package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", 3.14, Secret{"a", "b"}}
	r := reflect.ValueOf(A)
	fmt.Println(r.String())

	iType := r.Type()
	fmt.Println(iType)

	fmt.Printf("The %d fields of %s are:\n", iType.NumField(), iType)
	for j := 0; j < iType.NumField(); j++ {
		fmt.Printf("\t%s", iType.Field(j).Name)
		fmt.Printf("\twith type: %s", r.Field(j).Type())
		fmt.Printf("\twith value: %v\n", r.Field(j).Interface())

		k := reflect.TypeOf(r.Field(j).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(j).Type())
		}

		if k == reflect.Struct {
			fmt.Println(r.Field(j).Type())
		}
	}
}
