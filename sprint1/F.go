package main

import (
	"bufio"
	"fmt"
	"os"
)

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
