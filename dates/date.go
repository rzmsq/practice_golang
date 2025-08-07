package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("usage: dates parse_string")
		return
	}
	dateString := os.Args[1]

	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Hour:", d.Hour(), "Minute:", d.Minute())
	}

	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Hour:", d.Hour(), "Minute:", d.Minute())
	}

	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Hour:", d.Hour(), "Minute:", d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)

	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time %d:%d:%d\n", d.Hour(), d.Minute(), d.Second())
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
}
