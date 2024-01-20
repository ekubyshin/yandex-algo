package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Алла захотела, чтобы у неё под окном были узкие клумбы с тюльпанам. На схеме земельного участка клумбы обозначаются просто горизонтальными отрезками, лежащими на одной прямой. Для ландшафтных работ было нанято n садовников. Каждый из них обрабатывал какой-то отрезок на схеме. Процесс был организован не очень хорошо, иногда один и тот же отрезок или его часть могли быть обработаны сразу несколькими садовниками. Таким образом, отрезки, обрабатываемые двумя разными садовниками, сливаются в один. Непрерывный обработанный отрезок затем станет клумбой. Нужно определить границы будущих клумб.
Рассмотрим примеры.
Пример 1:
Даны 4 отрезка:
[7,8]
,
[7,8]
 ,
[2,3]
,
[6,10]
. Два одинаковых отрезка
[7,8]
 и
[7,8]
 сливаются в один, но потом их накрывает отрезок
[6,10]
. Таким образом, имеем две клумбы с координатами
[2,3]
 и
[6,10]
.
Пример 2
Во втором сэмпле опять 4 отрезка:
[2,3]
,
[3,4]
,
[3,4]
,
[5,6]
. Отрезки
[2,3]
,
[3,4]
 и
[3,4]
 сольются в один отрезок
[2,4]
. Отрезок
[5,6]
 ни с кем не объединяется, добавляем его в ответ.

Формат ввода
В первой строке задано количество садовников
n
. Число садовников не превосходит
100000.
В следующих
n
 строках через пробел записаны координаты клумб в формате: start end, где start —– координата начала, end —– координата конца. Оба числа целые, неотрицательные и не превосходят
107
. start строго меньше, чем end.

Формат вывода
Нужно вывести координаты каждой из получившихся клумб в отдельных строках. Данные должны выводиться в отсортированном порядке —– сначала клумбы с меньшими координатами, затем —– с бОльшими.
Пример 1
Ввод	Вывод
4
7 8
7 8
2 3
6 10
2 3
6 10

*/

func main() {
	var n int
	fmt.Scanf("%d", &n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	pairs := make([]Pair, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		t := scanner.Text()
		arrs := strings.Split(t, " ")
		l, _ := strconv.Atoi(arrs[0])
		r, _ := strconv.Atoi(arrs[1])
		pairs = append(pairs, NewPair(l, r))
	}
	res := solve(pairs)
	b := bufio.NewWriter(os.Stdout)
	for _, r := range res {
		b.WriteString(r.String())
		b.WriteString("\n")
	}
	b.Flush()
}

func solve(pairs []Pair) []Pair {
	if len(pairs) == 1 {
		return pairs
	}
	slices.SortStableFunc(pairs, compare)
	res := make([]Pair, 0, len(pairs))
	uniq := pairs[0]
	wasChanged := false
	for i := 1; i < len(pairs); i++ {
		if uniq.Right >= pairs[i].Right || uniq.Right >= pairs[i].Left {
			min := uniq.Left
			if min > pairs[i].Left {
				min = pairs[i].Left
			}
			max := uniq.Right
			if pairs[i].Right > max {
				max = pairs[i].Right
			}
			uniq = Pair{min, max}
			wasChanged = true
		} else {
			res = append(res, uniq)
			uniq = pairs[i]
			wasChanged = false
			if i == len(pairs)-1 {
				res = append(res, pairs[i])
			}
		}
	}
	if wasChanged {
		res = append(res, uniq)
	}
	return res
}

type Pair struct {
	Left  int
	Right int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d %d", p.Left, p.Right)
}

func (p Pair) Compare(p2 Pair) int {
	if p.Left < p2.Left {
		return -1
	}
	if p.Left > p2.Left {
		return 1
	}
	if p.Left == p2.Left {
		if p.Right > p2.Right {
			return 1
		}
		if p.Right < p2.Right {
			return -1
		}
	}
	return 0
}

func NewPair(l int, r int) Pair {
	return Pair{l, r}
}

func compare(p1 Pair, p2 Pair) int {
	return p1.Compare(p2)
}
