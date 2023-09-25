package main

//91290012
// proc O(N*M+9) N - кол-во строк M - кол-во столбцов
// mem O(9) ну по сути O(1) константа

import (
	"bufio"
	"errors"
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
	for i := 0; i < max+rows; i++ {
		if b, ok := reader.ReadByte(); ok == nil {
			num, err := getInt(b)
			if err == nil {
				m[num] += 1
			}
		} else {
			break
		}
	}
	return m
}

func getInt(b byte) (int, error) {
	if b == 10 {
		return 0, NewLineEnded()
	}
	if b >= '1' && b <= '9' {
		num, err := strconv.Atoi(string(b))
		if err == nil {
			return num, nil
		}
	}
	return 0, NewNotANumber()
}

type LineEnded error

type NotANumber error

func NewLineEnded() LineEnded {
	return errors.New("line is ended")
}

func NewNotANumber() error {
	return errors.New("not a number")
}
