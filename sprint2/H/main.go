package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Stack struct {
	items []byte
}

func NewStack(size int) Stack {
	return Stack{
		items: make([]byte, 0, size),
	}
}

func (s *Stack) push(b byte) {
	s.items = append(s.items, b)
}

func (s *Stack) pop() (byte, error) {
	if len(s.items) == 0 {
		return 0, errors.New("empty stack")
	}
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem, nil
}

func (s *Stack) size() int {
	return len(s.items)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([]byte, 0, 1000)
	for {
		if b, ok := reader.ReadByte(); ok == nil {
			if b == 10 {
				break
			}
			arr = append(arr, b)
		} else {
			break
		}
	}
	switch is_correct_bracket_seq(arr) {
	case true:
		fmt.Println("True")
	default:
		fmt.Println("False")
	}
}

func is_correct_bracket_seq(arr []byte) bool {
	stack := NewStack(len(arr))
	for _, b := range arr {
		if b == '{' || b == '[' || b == '(' {
			stack.push(b)
		} else {
			p, err := stack.pop()
			if err != nil {
				return false
			} else if !checkSkobka(b, p) {
				return false
			}
		}
	}
	return stack.size() == 0
}

func checkSkobka(close byte, open byte) bool {
	switch close {
	case ')':
		return open == '('
	case ']':
		return open == '['
	case '}':
		return open == '{'
	}
	return false
}
