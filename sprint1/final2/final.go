package main

//91290012
// proc O(N*M+9) N - кол-во строк M - кол-во столбцов
// mem O(9) ну по сути O(1) константа

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
Игра «Тренажёр для скоростной печати» представляет собой поле 4x4 из клавиш, на которых — либо точка, либо цифра от одного до девяти. Суть игры следующая: каждый раунд на поле появляется комбинация цифр и точек. В момент времени t игрок должен одновременно нажать на все клавиши, где есть цифра t.

Если в момент t нажаты все нужные клавиши, игроки получают один балл. Если клавиш с цифрой t на поле нет, победное очко не начисляется

Два игрока в один момент могут нажать на k клавиш каждый. Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём. Рассмотрим пример 1, в котором k=3./

Допустим, t=1. В таком случае один игрок должен нажать на две кнопки с цифрой 1. Чтобы узнать, сколько клавиш нажмут два игрока, воспользуемся формулой: k*2. Получается, что вместе мальчики нажмут на шесть клавиш и получат победное очко.

Когда t=2, двум игрокам необходимо нажать на семь кнопок одновременно. Но это не под силу ребятам: каждый может нажать только по три кнопки. Победное очко не начисляется.

При t=3, каждому игроку нужно нажать на одну кнопку. Успех! Теперь на счету Гоши и Тимофея целых два победных очка.

Других цифр на поле нет. Поэтому в следующих раундах, где t=4...t=9, победные очки начисляться не будут. Таким образом, Гоша и Тимофей заработают два балла.

Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём.

Формат ввода
В первой строке дано целое число k (1 ≤ k ≤ 5).

В четырёх следующих строках задан вид тренажёра -— по 4 символа в каждой строке. Каждый символ – либо точка, либо цифра от 1 до 9. Символы одной строки идут подряд и не разделены пробелами.

Формат вывода
Выведите единственное число -— максимальное количество баллов, которое смогут набрать Гоша и Тимофей.

Пример 1
Ввод	Вывод
3
1231
2..2
2..2
2..2
2
Пример 2
Ввод	Вывод
4
1111
9999
1111
9911
1
Пример 3
Ввод	Вывод
4
1111
1111
1111
1111

*/

const (
	rows = 4
	cols = 4
)

func main() {
	var k int
	fmt.Scanln(&k)
	m := readArray()
	out := 0
	max := k * 2
	for i := 1; i < len(m); i++ {
		if m[i] <= max && m[i] != 0 {
			out += 1
		}
	}
	fmt.Println(out)
}

func readArray() []int {
	reader := bufio.NewReader(os.Stdin)
	max := rows * cols
	m := make([]int, 10)
	for i := 0; i < max+rows; i++ {
		if b, ok := reader.ReadByte(); ok == nil {
			num, err := getInt(b)
			if err == nil {
				m[num] += 1
			}
		} else {
			break
		}
	}
	return m
}

func getInt(b byte) (int, error) {
	if b == 10 {
		return 0, NewLineEnded()
	}
	if b >= '1' && b <= '9' {
		num, err := strconv.Atoi(string(b))
		if err == nil {
			return num, nil
		}
	}
	return 0, NewNotANumber()
}

type LineEnded error

type NotANumber error

func NewLineEnded() LineEnded {
	return errors.New("line is ended")
}

func NewNotANumber() error {
	return errors.New("not a number")
}
