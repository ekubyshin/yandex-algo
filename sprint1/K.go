package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]string, n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}
	scanner.Scan()
	s := scanner.Text()
	b := strings.Split(s, "")
	maxLen := len(arr)
	minLen := len(b)
	maxN := &arr
	minN := &b
	if maxLen < len(b) {
		maxLen = len(b)
		minLen = len(arr)
		maxN = &b
		minN = &arr
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
		if sum >= 10 {
			p = 1
			sum = sum - 10
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
		if i != len(buf)-1 {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
}
