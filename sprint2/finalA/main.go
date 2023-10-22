package main

//93251213
//cpu O(1) - операция вставки занимает константное время, тк доступ осуществляется по индексу в массиве
// Если учесть обработку ввода, то cpu O(n)
// Если учесть еще работу Split, то будет O(n + m), где m - кол-во операций
//mem O(1) - дополнительная память требуется только для хранения указателей и размера очереди, те - константа
//Реализован на кольцевом буфере
//указатели на конец и начало высчитываются по модулю
//Для избавления от копипасты реализовал вспомогательные методы декораторы над основными методами дэка
//Для упрощения подсчета элементов добавил переменную счетчик - при добавлении инкрементируется на 1, при взятии элемента - уменьшается на 1
//методы возвращают ошибку при невозможности совершения операции, которая перехватывается и обрабатывается в основной программе
//при вставке в конец элемент добавляется а указатель перепрыгивает на след место
//при вставке в начало сначала высчитывается новое место куда вставить с переполнением по длинне очереди и взятием по модулю
// получается, что при начальной вставке (0 - 1 + n) % n = n - 1
// получается, что буфер используем в двух направлениях, таким образом имитируя работу двусвязанного списка
// ввод и вывод буферезованы, таким образом не тратится время cpu на ввод/вывод
// все делается в одном цикле

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (q *Queue) PushBack(x int) error {
	if q.size >= q.maxN {
		return errors.New("size exceeded")
	}
	q.queue[q.tail] = x
	q.tail = q.nextIndex(q.tail)
	q.size += 1
	return nil
}

func (q *Queue) PopBack() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	q.tail = q.prevIndex(q.tail)
	x := q.queue[q.tail]
	q.queue[q.tail] = 0
	q.size -= 1
	return x, nil
}

func (q *Queue) PushFront(x int) error {
	if q.size >= q.maxN {
		return errors.New("size exceeded")
	}
	q.head = q.prevIndex(q.head)
	q.queue[q.head] = x
	q.size += 1
	return nil
}

func (q *Queue) PopFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	x := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = q.nextIndex(q.head)
	q.size -= 1
	return x, nil
}

func (q *Queue) prevIndex(cur int) int {
	return (cur - 1 + q.maxN) % q.maxN
}

func (q *Queue) nextIndex(cur int) int {
	return (cur + 1) % q.maxN
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

// декоратор для взятия элемента
func handlePop(f func() (int, error)) string {
	if el, err := f(); err == nil {
		return strconv.Itoa(el)
	}
	return "error"
}

// декоратор для вставки элемента
func handlePush(n int, f func(int) error) string {
	if err := f(n); err != nil {
		return "error"
	}
	return ""
}

// обработчик команд
// сначала проверяем односоставные команды
// иначе получаем аргументы для вызова команды с аргументом
func handleCommand(cmd string, queue *Queue) string {
	switch cmd {
	case "pop_back":
		return handlePop(queue.PopBack)
	case "pop_front":
		return handlePop(queue.PopFront)
	}
	arr := strings.Split(cmd, " ")
	num := arr[1]
	n, _ := strconv.Atoi(num)
	cmd = arr[0]
	switch cmd {
	case "push_back":
		return handlePush(n, queue.PushBack)
	case "push_front":
		return handlePush(n, queue.PushFront)
	}
	return ""
}
