package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Гоша любит играть в игру «Подпоследовательность»: даны 2 строки, и нужно понять, является ли первая из них подпоследовательностью второй. Когда строки достаточно длинные, очень трудно получить ответ на этот вопрос, просто посмотрев на них. Помогите Гоше написать функцию, которая решает эту задачу.

Формат ввода
В первой строке записана строка s.

Во второй —- строка t.

Обе строки состоят из маленьких латинских букв, длины строк не превосходят 150000. Строки не могут быть пустыми.

Формат вывода
Выведите True, если s является подпоследовательностью t, иначе —– False.

1.
abc
ahbgdcu

True

2.
abcp
ahpc

False


*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	source := readSequence(reader)
	input := readSequence(reader)
	if isSequence(source, input) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func isSequence(source []byte, input []byte) bool {
	si := 0
	for i := 0; i < len(input); i++ {
		if si >= len(source) {
			break
		}
		if source[si] == input[i] {
			si += 1
		}
	}
	return si == len(source)
}

func readSequence(reader *bufio.Reader) []byte {
	input := make([]byte, 0, 1000)
	for {
		b, err := reader.ReadByte()
		if err != nil || b == 10 {
			break
		}
		input = append(input, b)
	}
	return input
}
