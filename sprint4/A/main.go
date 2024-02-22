package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
polynomail hash for string


*/

func main() {
	fmt.Println(strconv.FormatUint(hash(read()), 10))
}

func hash(str string, q uint64, m uint64) uint64 {
	acc := uint64(0)
	for _, s := range str {
		acc = (acc*q + uint64(s)) % m
	}
	return acc
}

func read() (string, uint64, uint64) {
	data, _ := os.ReadFile("./input.txt")
	dataArray := strings.Split(string(data), "\n")
	q, _ := strconv.ParseUint(dataArray[0], 10, 64)
	m, _ := strconv.ParseUint(dataArray[1], 10, 64)
	return dataArray[2], q, m
}
