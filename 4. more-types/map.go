package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)

	m["Diogo Felipe"] = 12
	fmt.Println(m["Diogo Felipe"])

	var m2 = map[string]int{
		"Pedro": 13,
		"Leo":   14,
	}
	fmt.Println(m2)

	m2["Pedro"] = 10
	fmt.Println(m2["Pedro"])

	delete(m, "Pedro")
	element, ok := m["Pedro"]
	fmt.Println(element, ok)
}
