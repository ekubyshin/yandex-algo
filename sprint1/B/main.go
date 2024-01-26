package main

import (
	"fmt"
	"log"
	"math"
)

/*
Представьте себе онлайн-игру для поездки в метро: игрок нажимает на кнопку, и на экране появляются три случайных числа. Если все три числа оказываются одной чётности, игрок выигрывает.

Напишите программу, которая по трём числам определяет, выиграл игрок или нет.

Формат ввода
В первой строке записаны три случайных целых числа a, b и c. Числа не превосходят 109 по модулю.

Формат вывода
Выведите «WIN», если игрок выиграл, и «FAIL» в противном случае.

Пример 1
Ввод	Вывод
1 2 -3
FAIL
Пример 2
Ввод	Вывод
7 11 7
WIN
Пример 3
Ввод	Вывод
6 -2 0
WIN
*/

func main() {
	var a, b, c int
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		log.Fatal(err)
	}
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	c = int(math.Abs(float64(c)))
	if a%2 > 0 && b%2 > 0 && c%2 > 0 || a%2 == 0 && b%2 == 0 && c%2 == 0 {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}