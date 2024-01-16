package L

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
Вася решил накопить денег на два одинаковых велосипеда — себе и сестре. У Васи есть копилка, в которую каждый день он может добавлять деньги (если, конечно, у него есть такая финансовая возможность). В процессе накопления Вася не вынимает деньги из копилки.

У вас есть информация о росте Васиных накоплений — сколько у Васи в копилке было денег в каждый из дней.

Ваша задача — по заданной стоимости велосипеда определить

первый день, в которой Вася смог бы купить один велосипед,
и первый день, в который Вася смог бы купить два велосипеда.
Подсказка: решение должно работать за O(log n).

Формат ввода
В первой строке дано число дней n, по которым велись наблюдения за Васиными накоплениями. 1 ≤ n ≤ 106.

В следующей строке записаны n целых неотрицательных чисел. Числа идут в порядке неубывания. Каждое из чисел не превосходит 106.

В третьей строке записано целое положительное число s — стоимость велосипеда. Это число не превосходит 106.

Формат вывода
Нужно вывести два числа — номера дней по условию задачи.

Если необходимой суммы в копилке не нашлось, нужно вернуть -1 вместо номера дня.

1.
6
1 2 4 4 6 8
3
=
3 5


2.
6
1 2 4 4 4 4
3
=
3 -1

3.
6
1 2 4 4 4 4
10
=
-1 -1

*/

func main() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	vals := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		if v, ok := strconv.ParseInt(s, 10, 64); ok == nil {
			vals = append(vals, v)
		}
	}
	scanner.Scan()
	k := scanner.Text()
	var ki int64
	var ok error
	if ki, ok = strconv.ParseInt(k, 10, 64); ok != nil {
		fmt.Println(-1)
		return
	}
	a, b := solv(ki, vals)
	fmt.Println(a, b)
}

func solv(p int64, elems []int64) (int, int) {
	from := 0
	a := solver(p, elems, from, len(elems))
	if a != -1 {
		from = a
	}
	b := solver(p*2, elems, from, len(elems))
	if a != -1 {
		a += 1
	}
	if b != -1 {
		b += 1
	}
	return a, b
}

func solver(x int64, elems []int64, left, right int) int {
	if right <= left {
		if left < len(elems) && elems[left] >= x {
			return left
		} else if right < len(elems) && elems[right] >= x {
			return right
		}
		return -1
	}
	mid := div(left+right, 2)
	if elems[mid] >= x {
		return solver(x, elems, left, mid)
	}
	return solver(x, elems, mid+1, right)
}

func div(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
