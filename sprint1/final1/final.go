package main

//91744373
// proc O(2*N)
// mem O(N) если все участки пустые

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	scanner := makeScanner()
	arr, pos := readArray(scanner, n)
	printArray(solver(n, arr, pos))
}

func solver(n int, arr []int, pos []int) []int {
	prev := 0
	next := 1
	if len(pos) == 1 {
		next = 0
	}
	for i := 0; i < len(arr); i++ {
		if i > pos[prev] && i < pos[next] {
			arr[i] = min(i-pos[prev], pos[next]-i)
		} else if i > pos[next] {
			arr[i] = i - pos[next]
		} else if i < pos[prev] {
			arr[i] = pos[prev] - i
		} else if i == pos[next] {
			arr[i] = 0
			prev = next
			if next+1 < len(pos) {
				next += 1
			}
		}
	}
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner, n int) (arr []int, pos []int) {
	scanner.Split(bufio.ScanWords)
	arr = make([]int, n)
	pos = make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		arr[i] = num
		if num == 0 {
			pos = append(pos, i)
		}
	}
	return arr, pos
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
