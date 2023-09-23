package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	const num = 4
	cur := num
	if n == 1 || n == num {
		fmt.Println("True")
		return
	}
	if n%2 > 0 || n < num {
		fmt.Println("False")
		return
	}
	for {
		if cur >= n {
			break
		}
		cur *= num
	}
	if cur == n {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
