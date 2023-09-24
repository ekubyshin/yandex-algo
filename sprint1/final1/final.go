package main

//91266065
// proc O(2*N)
// mem O(N) если все участки пустые

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type cursor struct {
	pos  []int
	iter int
}

func makeCursor(arr []int) cursor {
	return cursor{pos: arr, iter: 0}
}

func (p *cursor) cur() int {
	return p.pos[p.iter]
}

func (p *cursor) hasNext() bool {
	return p.iter < len(p.pos)-1
}

func (p *cursor) next() (int, error) {
	if p.hasNext() {
		return p.pos[p.iter+1], nil
	}
	return 0, errors.New("no elem")
}

func (p *cursor) move() (int, error) {
	if p.hasNext() {
		p.iter += 1
		return p.pos[p.iter], nil
	}
	return 0, errors.New("no elem")
}

func main() {
	var n int
	fmt.Scanln(&n)
	scanner := makeScanner()
	arr, pos := readArray(scanner, n)
	printArray(solver(n, arr, pos))
}

func solver(n int, arr []int, pos []int) []int {
	for i, cur := range pos {
		next := len(arr)
		middle := next
		if i < len(pos)-1 {
			next = pos[i+1]
			middle = int(math.Ceil(float64(next-cur)/2.0)) + cur - 1
		}
		l := 0
		if i == 0 && pos[0] > 0 {
			for j := cur - 1; j >= 0; j-- {
				arr[j] = cur - j
			}
		}
		if cur < next {
			for j := cur + 1; j < next; j++ {
				if j <= middle {
					l += 1
					arr[j] = l
				} else {
					arr[j] = next - j
				}
			}
		}
	}
	return arr
}

func oldSolver(n int, arr []int, pos []int) []int {
	out := make([]int, n)
	cur := makeCursor(pos)
	p := 0
	for i, v := range arr {
		if v == 0 {
			out[i] = 0
			if cur.hasNext() {
				next, _ := cur.next()
				if i == next {
					cur.move()
				}
			}
			p = 0
			continue
		}
		if i <= cur.cur() || cur.hasNext() {
			if i < cur.cur() {
				out[i] = cur.cur() - i
				continue
			}
			next, _ := cur.next()
			middle := int(math.Ceil(float64(next-cur.cur())/2.0)) + cur.cur() - 1
			if i <= middle {
				p += 1
				out[i] = p
			} else {
				out[i] = next - i
			}
			if i == next || i <= cur.cur() {
				cur.move()
				p = 0
			}
		} else {
			p += 1
			out[i] = p
		}
	}
	return out
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
