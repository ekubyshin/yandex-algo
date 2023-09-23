package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	a = a[:]
	b = b[:]
	maxLen := len(a)
	minLen := len(b)
	maxN := &a
	minN := &b
	if maxLen < len(b) {
		maxLen = len(b)
		minLen = len(a)
		maxN = &b
		minN = &a
	}
	j := minLen - 1
	p := 0
	buf := make([]int, maxLen+1)
	for i := maxLen - 1; i >= 0; i-- {
		ai := 0
		bj := 0
		ai, _ = strconv.Atoi(string((*maxN)[i]))
		if j >= 0 {
			bj, _ = strconv.Atoi(string((*minN)[j]))
			j -= 1
		} else {
			bj = 0
		}
		sum := ai + bj + p
		if sum == 3 {
			sum = 1
			p = 1
		} else if sum == 2 {
			sum = 0
			p = 1
		} else {
			p = 0
		}
		buf[i+1] = sum
	}
	buf[0] = p
	writer := bufio.NewWriter(os.Stdout)
	for i, v := range buf {
		if i == 0 && v == 0 {
			continue
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.Flush()
}
