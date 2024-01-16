package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Астрологи объявили день очередей ограниченного размера. Тимофею нужно написать класс MyQueueSized, который принимает параметр max_size, означающий максимально допустимое количество элементов в очереди.

Помогите ему —– реализуйте программу, которая будет эмулировать работу такой очереди. Функции, которые надо поддержать, описаны в формате ввода.

Формат ввода
В первой строке записано одно число — количество команд, оно не превосходит 5000.
Во второй строке задан максимально допустимый размер очереди, он не превосходит 5000.
Далее идут команды по одной на строке. Команды могут быть следующих видов:

push(x) — добавить число x в очередь;
pop() — удалить число из очереди и вывести на печать;
peek() — напечатать первое число в очереди;
size() — вернуть размер очереди;
При превышении допустимого размера очереди нужно вывести «error». При вызове операций pop() или peek() для пустой очереди нужно вывести «None».
Формат вывода
Напечатайте результаты выполнения нужных команд, по одному на строке.

Пример 1
Ввод	Вывод
8
2
peek
push 5
push 2
peek
size
size
push 1
size
None
5
2
2
error
2
Пример 2
Ввод	Вывод
10
1
push 1
size
push 3
size
push 1
pop
push 1
pop
push 3
push 3
*/

type Queue struct {
	queue []int
	maxN  int
	head  int
	tail  int
	size  int
}

func NewQueue(n int) *Queue {
	return &Queue{
		queue: make([]int, n),
		maxN:  n,
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(x int) error {
	if q.size != q.maxN {
		q.queue[q.tail] = x
		q.tail = (q.tail + 1) % q.maxN
		q.size += 1
		return nil
	} else {
		return errors.New("size exceeded")
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.queue[q.head], nil
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	x := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.maxN
	q.size -= 1
	return x, nil
}

func main() {
	var cmds, size int
	fmt.Scanln(&cmds)
	fmt.Scanln(&size)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	queue := NewQueue(size)
	writer := bufio.NewWriter(os.Stdout)
	for ; cmds > 0; cmds-- {
		scanner.Scan()
		cmd := scanner.Text()
		if s := handleCommand(cmd, queue); s != "" {
			writer.WriteString(s)
			if cmds > 1 {
				writer.WriteString("\n")
			}
		}
	}
	writer.Flush()
}

func handleCommand(cmd string, queue *Queue) string {
	if cmd == "peek" {
		if el, err := queue.Peek(); err == nil {
			return strconv.Itoa(el)
		} else {
			return "None"
		}
	} else if cmd == "size" {
		return strconv.Itoa(queue.Size())
	} else if cmd == "pop" {
		if el, err := queue.Pop(); err == nil {
			return strconv.Itoa(el)
		} else {
			return "None"
		}
	}
	num := strings.Split(cmd, " ")[1]
	n, _ := strconv.Atoi(num)
	if err := queue.Push(n); err != nil {
		return "error"
	}
	return ""
}
