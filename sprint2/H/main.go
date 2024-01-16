package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/*
Вот какую задачу Тимофей предложил на собеседовании одному из кандидатов. Если вы с ней ещё не сталкивались, то наверняка столкнётесь –— она довольно популярная.

Дана скобочная последовательность. Нужно определить, правильная ли она.

Будем придерживаться такого определения:

пустая строка —– правильная скобочная последовательность;
правильная скобочная последовательность, взятая в скобки одного типа, –— правильная скобочная последовательность;
правильная скобочная последовательность с приписанной слева или справа правильной скобочной последовательностью —– тоже правильная.
На вход подаётся последовательность из скобок трёх видов: [], (), {}.
Напишите функцию is_correct_bracket_seq, которая принимает на вход скобочную последовательность и возвращает True, если последовательность правильная, а иначе False.

Формат ввода
На вход подаётся одна строка, содержащая скобочную последовательность. Скобки записаны подряд, без пробелов.

Формат вывода
Выведите «True» или «False».

Пример 1
Ввод	Вывод
{[()]}
True
Пример 2
Ввод	Вывод
()
True
*/

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
