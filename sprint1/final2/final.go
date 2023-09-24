package main

//91221240
// proc O(N*M+9) N - кол-во строк M - кол-во столбцов
// mem O(9) ну по сути O(1) константа

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	rows = 4
	cols = 4
)

func main() {
	var k int
	fmt.Scanln(&k)
	m := readArray()
	out := 0
	max := k * 2
	for i := 1; i < len(m); i++ {
		if m[i] <= max && m[i] != 0 {
			out += 1
		}
	}
	fmt.Println(out)
}

func readArray() []int {
	reader := bufio.NewReader(os.Stdin)
	max := rows * cols
	m := make([]int, 10)
	i := 0
	for i < max {
		if b, ok := reader.ReadByte(); ok == nil {
			if b == 10 {
				continue
			}
			num := 0
			if b >= '1' && b <= '9' {
				num, _ = strconv.Atoi(string(b))
				m[num] += 1
			}
			i += 1
		} else {
			break
		}
	}
	return m
}
