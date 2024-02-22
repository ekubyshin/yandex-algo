package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Жители Алгосского архипелага придумали новый способ сравнения строк. Две строки считаются равными, если символы одной из них можно заменить на символы другой так, что первая строка станет точной копией второй строки. При этом необходимо соблюдение двух условий:

Порядок вхождения символов должен быть сохранён.
Одинаковым символам первой строки должны соответствовать одинаковые символы второй строки. Разным символам —– разные.
Например, если строка s = «abacaba», то ей будет равна строка t = «xhxixhx», так как все вхождения «a» заменены на «x», «b» –— на «h», а «c» –— на «i». Если же первая строка s=«abc», а вторая t=«aaa», то строки уже не будут равны, так как разные буквы первой строки соответствуют одинаковым буквам второй.

алгоритм решения
сравнивать буду не конкретные символы, а паттерны обоих строк.
Для каждой строки считаю паттерн символов вида - символ = количество, с сохранением порядка
Для этого сделаю реализацию OrderedMap из задания D
Потом просто бегу по циклу, до тех пор, пока паттерны различаютися и прерываю цикл сравнения паттернов

В качестве усовершенствования по памяти, можно скомбинировать паттерн как комбинацию rune1, rune2 и count, где count будет инкрементиться из str1 и декриментиться из str2
*/

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

func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	if !o.Contains(key) {
		return *(new(V)), false
	}
	return o.list[o.keys[key]].Value, true
}

func (o *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := o.keys[key]
	return ok
}

func (o *OrderedMap[K, V]) Update(key K, value V) bool {
	if !o.Contains(key) {
		o.Put(key, value)
		return false
	}
	o.list[o.keys[key]].Value = value
	return true
}

func (o *OrderedMap[K, V]) Pairs() []Pair[K, V] {
	return o.list[:]
}

func main() {
	if solve(read()) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(str1 string, str2 string) bool {
	pattern1 := makePattern(str1).Pairs()
	pattern2 := makePattern(str2).Pairs()
	for i, v := range pattern1 {
		if pattern2[i].Value != v.Value {
			return false
		}
	}
	return true
}

func makePattern(str string) *OrderedMap[rune, int] {
	pattern := NewOrderedMap[rune, int](len(str))
	for _, s := range str {
		if c, ok := pattern.Get(s); ok {
			pattern.Update(s, c+1)
		} else {
			pattern.Put(s, 1)
		}
	}
	return pattern
}

func read() (string, string) {
	data, _ := os.ReadFile("./input.txt")
	dataArray := strings.Split(string(data), "\n")
	return dataArray[0], dataArray[1]
}
