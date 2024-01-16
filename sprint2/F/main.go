package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Нужно реализовать класс StackMax, который поддерживает операцию определения максимума среди всех элементов в стеке. Класс должен поддерживать операции push(x), где x – целое число, pop() и get_max().

Формат ввода
В первой строке записано одно число n — количество команд, которое не превосходит 10000. В следующих n строках идут команды. Команды могут быть следующих видов:

push(x) — добавить число x в стек. Число x не превышает 105;
pop() — удалить число с вершины стека;
get_max() — напечатать максимальное число в стеке;
Если стек пуст, при вызове команды get_max() нужно напечатать «None», для команды pop() — «error».

Формат вывода
Для каждой команды get_max() напечатайте результат её выполнения. Если стек пустой, для команды get_max() напечатайте «None». Если происходит удаление из пустого стека — напечатайте «error».

Пример 1
Ввод	Вывод
8
get_max
push 7
pop
push -2
push -1
pop
get_max
get_max
None
-2
-2
Пример 2
Ввод	Вывод
7
get_max
pop
pop
pop
push 10
get_max
push -9
*/

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
