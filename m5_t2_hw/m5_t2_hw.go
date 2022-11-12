/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	var hashStr string
	var resultStr string
	fmt.Println("Введите строку для расшифровки")
	_, err := fmt.Scanf("%s", &hashStr)
	if err != nil {
		log.Fatal("Epic fail: " + err.Error())
	}

	letterMap := fillLetterMap()
	for i := 0; i < len(hashStr); i += 2 {
		substr := string(hashStr[i]) + string(hashStr[i+1])
		k, ok := letterMap[substr]
		if !ok {
			fmt.Println("не существует символа для кода", substr)
			return
		}

		resultStr += k
	}

	fmt.Println("Расшифровка строки =", resultStr)
}

func fillLetterMap() map[string]string {
	var letters map[string]string
	letters = make(map[string]string)

	i := 0
	for asciiNum := 97; asciiNum < 123; asciiNum++ {
		character := string(asciiNum)
		index := fmt.Sprintf("%02d", i)
		letters[index] = character
		i++
	}
	letters["26"] = " "
	return letters
}
