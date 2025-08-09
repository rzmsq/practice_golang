package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	sumTrue := 0
	sumFalse := 0
	for i, arg := range arguments {
		if i == 0 {
			continue
		}
		if matchInt(arg) {
			sumTrue++
		} else {
			sumFalse++
		}
	}
	fmt.Println(sumTrue, sumFalse)

}
