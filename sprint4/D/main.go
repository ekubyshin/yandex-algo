package main

/*
В компании, где работает Тимофей, заботятся о досуге сотрудников и устраивают различные кружки по интересам. Когда кто-то записывается на занятие, в лог вносится название кружка.

По записям в логе составьте список всех кружков, в которые ходит хотя бы один человек.

Формат ввода
В первой строке даётся натуральное число n, не превосходящее 10 000 –— количество записей в логе.

В следующих n строках —– названия кружков.

Формат вывода
Выведите уникальные названия кружков по одному на строке, в порядке появления во входных данных.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func NewOrderedMap[K comparable, V any](size int) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys: make(map[K]int, size),
		list: make([]Pair[K, V], 0, size),
	}
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type OrderedMap[K comparable, V any] struct {
	keys     map[K]int
	list     []Pair[K, V]
	curIndex int
}

func (o *OrderedMap[K, V]) Put(key K, value V) {
	p := Pair[K, V]{key, value}
	if v, ok := o.keys[key]; ok {
		o.list[v] = p
		return
	}
	o.keys[key] = o.curIndex
	o.list = append(o.list, p)
	o.curIndex += 1
}

func (o *OrderedMap[K, V]) Pairs() []Pair[K, V] {
	return o.list[:]
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	sections := NewOrderedMap[string, struct{}](num)
	for i := 0; i < num; i++ {
		scanner.Scan()
		name := scanner.Text()
		sections.Put(name, struct{}{})
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range sections.Pairs() {
		writer.WriteString(v.Key)
		writer.WriteString("\n")
	}
	writer.Flush()
}
