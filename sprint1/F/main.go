package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Помогите Васе понять, будет ли фраза палиндромом‎. Учитываются только буквы и цифры, заглавные и строчные буквы считаются одинаковыми.

Решение должно работать за O(N), где N — длина строки на входе.

Формат ввода
В единственной строке записана фраза или слово. Буквы могут быть только латинские. Длина текста не превосходит 20000 символов.

Фраза может состоять из строчных и прописных латинских букв, цифр, знаков препинания.

Формат вывода
Выведите «True», если фраза является палиндромом, и «False», если не является.

Пример 1
Ввод	Вывод
A man, a plan, a canal: Panama
True
Пример 2
Ввод	Вывод
zo
False
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	letters := make([]string, 0, 1000)
	for {
		if b, ok := reader.ReadByte(); ok == nil {
			if b == 10 {
				break
			}
			if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
				if 'A' <= b && b <= 'Z' {
					b += 'a' - 'A'
				}
				letters = append(letters, string(b))
			}
		} else {
			break
		}
	}
	i := 0
	j := len(letters) - 1
	isPolindrom := true
	for {
		if i < j {
			if isPolindrom && letters[i] == letters[j] {
				i += 1
				j -= 1
			} else {
				isPolindrom = false
				break
			}
		} else {
			break
		}
	}
	if isPolindrom {
		fmt.Println("True")
		return
	}
	fmt.Println("False")
}
