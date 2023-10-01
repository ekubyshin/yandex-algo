package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MaxStack struct {
	max  []int
	data []int
}

func NewMaxStack(size int) MaxStack {
	return MaxStack{
		max:  make([]int, 0, size),
		data: make([]int, 0, size),
	}
}

func (s *MaxStack) push(x int) {
	if len(s.data) == 0 || s.max[len(s.max)-1] < x {
		s.max = append(s.max, x)
	} else {
		s.max = append(s.max, s.max[len(s.max)-1])
	}
	s.data = append(s.data, x)
}

func (s *MaxStack) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
		return
	}
	s.max = s.max[:len(s.max)-1]
	s.data = s.data[:len(s.data)-1]
}

func (s *MaxStack) get_max() {
	if len(s.data) == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(s.max[len(s.max)-1])
}

func main() {
	var size int
	fmt.Scanln(&size)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	maxStack := NewMaxStack(size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		command := scanner.Text()
		if command == "get_max" {
			maxStack.get_max()
			continue
		}
		if command == "pop" {
			maxStack.pop()
			continue
		}
		cmds := strings.Split(command, " ")
		n, _ := strconv.Atoi(cmds[1])
		maxStack.push(n)
	}
}
