package main

import "fmt"

func main() {
	arr := [...]string{"Ivan", "Vasya", "Petya", "Sasha", "Masha", "Dasha", "Vanya", "Kolya", "Olya", "Lena"}

	m := make(map[string]int)

	for _, v := range arr {
		if _, exists := m[v]; !exists {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	fmt.Println(m)

	keys := make([]string, 0, len(m))
	values := make([]int, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	fmt.Println(keys, values)

}
