package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

/*
Задача:
Тимофей решил организовать соревнование по спортивному программированию, чтобы найти талантливых стажёров. Задачи подобраны, участники зарегистрированы, тесты написаны. Осталось придумать, как в конце соревнования будет определяться победитель.

Каждый участник имеет уникальный логин. Когда соревнование закончится, к нему будут привязаны два показателя: количество решённых задач Pi и размер штрафа Fi. Штраф начисляется за неудачные попытки и время, затраченное на задачу.

Тимофей решил сортировать таблицу результатов следующим образом: при сравнении двух участников выше будет идти тот, у которого решено больше задач. При равенстве числа решённых задач первым идёт участник с меньшим штрафом. Если же и штрафы совпадают, то первым будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке.

Тимофей заказал толстовки для победителей и накануне поехал за ними в магазин. В своё отсутствие он поручил вам реализовать алгоритм быстрой сортировки (англ. quick sort) для таблицы результатов. Так как Тимофей любит спортивное программирование и не любит зря расходовать оперативную память, то ваша реализация сортировки не может потреблять O(n) дополнительной памяти для промежуточных данных (такая модификация быстрой сортировки называется "in-place").

Как работает in-place quick sort
Как и в случае обычной быстрой сортировки, которая использует дополнительную память, необходимо выбрать опорный элемент (англ. pivot), а затем переупорядочить массив. Сделаем так, чтобы сначала шли элементы, не превосходящие опорного, а затем —– большие опорного.

Затем сортировка вызывается рекурсивно для двух полученных частей. Именно на этапе разделения элементов на группы в обычном алгоритме используется дополнительная память. Теперь разберёмся, как реализовать этот шаг in-place.

Пусть мы как-то выбрали опорный элемент. Заведём два указателя left и right, которые изначально будут указывать на левый и правый концы отрезка соответственно. Затем будем двигать левый указатель вправо до тех пор, пока он указывает на элемент, меньший опорного. Аналогично двигаем правый указатель влево, пока он стоит на элементе, превосходящем опорный. В итоге окажется, что что левее от left все элементы точно принадлежат первой группе, а правее от right — второй. Элементы, на которых стоят указатели, нарушают порядок. Поменяем их местами (в большинстве языков программирования используется функция swap()) и продвинем указатели на следующие элементы. Будем повторять это действие до тех пор, пока left и right не столкнутся.
На рисунке представлен пример разделения при pivot=5. Указатель left — голубой, right — оранжевый.

Решение:
CPU O(nlogn)
Mem O(1)

Реализую структуру Competitor, которая будет содержать три поля Имя участника, его очки и размер штрафа
Далее для Competitor реализую метод Compare, который будет сравнивать одного участника с другим.
Добавлю дополнительный метод Compare, который сравнивает двух участников между собой в нужном порядке с возрастания на убывание
Для красоты кода, добавлю алиас типа на функцию сравнения, которая будет передаваться в метод quicksort
Данная функция сравнения, будет использоваться для расстановки участников слева или справа
Также добавлю два вспомогательных метода readCompetitors - читает из stdin и возвращает массив участников
writeResults - выводит результаты в stdout. Ввод и вывод буферизованные.
*/

type Competitor struct {
	Name    []byte
	Score   int
	Penalty int
}

func (c Competitor) Compare(b Competitor) int {
	if c.Score > b.Score {
		return 1
	}
	if c.Score == b.Score {
		if c.Penalty < b.Penalty {
			return 1
		}
		if c.Penalty == b.Penalty {
			return bytes.Compare(b.Name, c.Name)
		}
	}
	return -1
}

func Compare(a Competitor, b Competitor) int {
	return b.Compare(a)
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	reader := bufio.NewReader(os.Stdin)
	competitors := readCompetitors(reader, num)
	quicksort(competitors, Compare)
	printResult(competitors)
}

func readCompetitors(reader *bufio.Reader, num int) []Competitor {
	competitors := make([]Competitor, num)
	for i := 0; i < num; i++ {
		var name []byte
		var score int
		var penalty int
		fmt.Scanf("%s %d %d", &name, &score, &penalty)
		competitors[i] = Competitor{
			name,
			score,
			penalty,
		}
	}
	return competitors
}

func printResult(competitors []Competitor) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(competitors); i++ {
		writer.Write(competitors[i].Name)
		if i != len(competitors)-1 {
			writer.WriteString("\n")
		}
	}
	writer.Flush()
}

func quicksort[S ~[]E, E any](array S, cmp func(a, b E) int) {
	if len(array) < 2 {
		return
	}
	inplaceQuickSort(array, 0, len(array)-1, cmp)
}

func inplaceQuickSort[S ~[]E, E any](array S, left int, right int, cmp func(a, b E) int) {
	if left >= right {
		return
	}
	start := left
	end := right
	// pivot := array[rand.Intn(len(array))]
	pivot := array[(start+end)/2]
	for start <= end {
		for cmp(array[start], pivot) == -1 {
			start += 1
		}
		for cmp(array[end], pivot) == 1 {
			end -= 1
		}
		if start <= end {
			swap(array, start, end)
			start += 1
			end -= 1
		}
	}
	inplaceQuickSort(array, left, end, cmp)
	inplaceQuickSort(array, start, right, cmp)
}

func swap[S ~[]E, E any](array S, left int, right int) {
	tmp := array[left]
	array[left] = array[right]
	array[right] = tmp
}
