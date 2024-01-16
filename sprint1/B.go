package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var a, b, c int
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		log.Fatal(err)
	}
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	c = int(math.Abs(float64(c)))
	if a%2 > 0 && b%2 > 0 && c%2 > 0 || a%2 == 0 && b%2 == 0 && c%2 == 0 {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}
