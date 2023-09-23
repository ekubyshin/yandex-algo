package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		w := scanner.Text()
		arr[i], _ = strconv.Atoi(string(w))
	}
	if n == 1 {
		fmt.Println(1)
		return
	}
	c := 0
	for i := 0; i < n; i++ {
		points := 0
		hasPrev := i > 0
		hasNext := i < n-1
		if !hasPrev || !hasNext {
			points += 1
		}
		if hasPrev && arr[i] > arr[i-1] {
			points += 1
		}
		if hasNext && arr[i] > arr[i+1] {
			points += 1
		}
		if points == 2 {
			c += 1
		}
	}
	fmt.Println(c)
}
