package main

//91336495
// proc O(2*N)
// mem O(N) если все участки пустые

import (
	"bufio"
	"fmt"
	"math"
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
	first := pos[0]
	if first > 0 {
		for i := first - 1; i >= 0; i-- {
			arr[i] = first - i
		}
	}

	last := pos[len(pos)-1]
	if last < len(arr)-1 {
		for i := last + 1; i < len(arr); i++ {
			arr[i] = i - last
		}
	}

	for i, cur := range pos {
		next := last
		middle := next
		if i < len(pos)-1 {
			next = pos[i+1]
			middle = int(math.Ceil(float64(next-cur)/2.0)) + cur - 1
		}
		l := 0
		if cur >= next {
			continue
		}
		for j := cur + 1; j < next; j++ {
			if j <= middle {
				l += 1
				arr[j] = l
			} else {
				arr[j] = next - j
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
