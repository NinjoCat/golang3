/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	digits    = "0123456789"
	lowercase = "abcdefghiklmnopqrstuvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTUVXYZ"
	special   = "_!@#$%^&"
)

func main() {
	var password string

	digitsMap := fillLettersMap(digits)
	lowercaseMap := fillLettersMap(lowercase)
	uppercaseMap := fillLettersMap(uppercase)
	specialMap := fillLettersMap(special)

	for i := 1; i < 6; i++ {
		fmt.Println("Введите пароль")
		_, err := fmt.Scanf("%s", &password)
		if err != nil {
			log.Fatal("Epic fail: " + err.Error())
		}

		_, errorDescription, passErr := checkPassword(password, digitsMap, lowercaseMap, uppercaseMap, specialMap)

		fmt.Println("Попытка номер ", i)

		if passErr != nil {
			fmt.Println("Ошибка : " + passErr.Error())
			for _, errorStr := range errorDescription {
				fmt.Println(errorStr)
			}
		} else {
			fmt.Println("Хороший пароль : " + password)
			break
		}
	}
}

func checkPassword(password string, digitsMap, lowercaseMap, uppercaseMap, specialMap map[string]string) (bool, []string, error) {
	var digitsExists bool
	var lowercaseExists bool
	var uppercaseExists bool
	var speciaExists bool
	var errorsDescriptions []string
	errorsDescriptions = make([]string, 0)

	if len(password) < 8 && len(password) > 15 {
		errorsDescriptions = append(errorsDescriptions, "Пароль должен содержать от 8 до 15 символов")
	}

	for _, k := range password {
		str := string(k)

		_, digitsMapCheck := digitsMap[str]
		_, lowercaseMapCheck := lowercaseMap[str]
		_, uppercaseMapCheck := uppercaseMap[str]
		_, specialMapCheck := specialMap[str]

		if digitsMapCheck {
			digitsExists = true
		}

		if lowercaseMapCheck {
			lowercaseExists = true
		}

		if uppercaseMapCheck {
			uppercaseExists = true
		}

		if specialMapCheck {
			speciaExists = true
		}

		if !digitsMapCheck && !lowercaseMapCheck && !uppercaseMapCheck && !specialMapCheck {
			errorsDescriptions = append(errorsDescriptions, "Недопустимый символ : "+str)
		}
	}

	if !digitsExists {
		errorsDescriptions = append(errorsDescriptions, "Не хватает цифр")
	}

	if !lowercaseExists {
		errorsDescriptions = append(errorsDescriptions, "Не хватает букв в нижнем регистре")
	}

	if !uppercaseExists {
		errorsDescriptions = append(errorsDescriptions, "Не хватает букв в верхнем регистре")
	}

	if !speciaExists {
		errorsDescriptions = append(errorsDescriptions, "Не хватает спец символов")
	}

	if len(errorsDescriptions) > 0 {
		return false, errorsDescriptions, errors.New("Плохой пароль")
	}
	return true, errorsDescriptions, nil
}

func fillLettersMap(str string) map[string]string {
	var letters map[string]string
	letters = make(map[string]string)

	for _, k := range str {
		letters[string(k)] = string(k)
	}

	return letters
}
