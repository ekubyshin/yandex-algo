package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//93250629
//cpu O(1) судя по условиям задачи, стэк не будет превышать 3, соответственно доступ к элементам O(1)
// Если учитывать разбитие строки по символам - то O(n), где n - кол-во операндов и операторов
//mem O(1) константа, тк доп память не требуется. Если, конечно возникнет ситуация,
// когда будут перечислены сначала все операнды, а потом кучей операторы, то стэк разрастется до O(n/2) что есть O(n)
// принцип работы: читается в буфер строка и разбивается на слова
// если это оператор, то определяется по таблице какой оператор вызывать
// если это не оператор, то значит операнд и кладется на стэк, до момента вызова оператора
// функция целочисленного деления определяется как округление по верхней границе деления дробей
// кажется, что можно оптимизировать работу и отказаться от приведения типо, дабы ускорить работу,
// тк с float операции в принципе занимают больше тактов, чем с инт, но это не является узким местом и можно принебречь

type Stack struct {
	data []int
}

func NewStack() Stack {
	return Stack{
		data: make([]int, 0, 10),
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("no elemets")
	}
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return el, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	stack := NewStack()
	scanner.Scan()
	line := scanner.Text()
	vals := strings.Split(line, " ")
	for _, v := range vals {
		handleInput(string(v), &stack)
	}
	res, _ := stack.Pop()
	fmt.Println(res)
}

type calcAction func(int, int) int

var calcActions = map[string]calcAction{
	"-": minus,
	"+": plus,
	"*": mult,
	"/": div,
}

func determineAction(inp string) (*calcAction, error) {
	if act, ok := calcActions[inp]; ok {
		return &act, nil
	}
	return nil, errors.New("not found action")
}

func handleInput(inp string, stack *Stack) {
	act, err := determineAction(inp)
	if err != nil {
		if num, err := strconv.Atoi(inp); err == nil {
			stack.Push(num)
		}
		return
	}
	lst, _ := stack.Pop()
	prev, _ := stack.Pop()
	stack.Push((*act)(prev, lst))
}

func minus(a, b int) int {
	return a - b
}

func plus(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
