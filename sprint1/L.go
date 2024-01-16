package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	am := make(map[rune]int)
	var unknown rune
	for _, c := range a {
		if v, ok := am[c]; ok {
			am[c] = v + 1
		} else {
			am[c] = 1
		}
	}
	for _, c := range b {
		if v, ok := am[c]; ok {
			if v == 0 {
				unknown = c
				break
			}
			am[c] -= 1
		} else {
			unknown = c
			break
		}
	}
	fmt.Println(string(unknown))
}
