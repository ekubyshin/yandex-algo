package main

import (
	"fmt"
	"log"
)

func main() {
	var a, x, b, c int
	_, err := fmt.Scanln(&a, &x, &b, &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(a*x*x + b*x + c)
}
