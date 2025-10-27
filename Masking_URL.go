package main

import (
	"bufio"
	"fmt"
	"os"
)

// функция для форматирования ввода
func receiveUserInput(userInput string) string {
	pattern := "http://"
	patternBytes := []byte(pattern)
	userInputSlice := []byte(userInput)
	result := make([]byte, 0, len(userInput)*2)
	currentPos := 0

	// Ищем вхождения паттерна(http://) в слайсе ввода
	for i := 0; i <= len(userInput)-len(patternBytes); i++ {
		match := true
		for j := 0; j < len(patternBytes); j++ {
			if userInputSlice[i+j] != patternBytes[j] {
				match = false
				break
			}
		}
		// Если паттерн найден
		if match {
			// Добавляем часть текста до найденного паттерна
			result = append(result, userInputSlice[currentPos:i]...)
			// Добавляем паттерн (http://)
			result = append(result, patternBytes...)

			// Начало данных после паттерна
			startIndex := i + len(patternBytes)
			remainingSlice := userInputSlice[startIndex:]

			// Ищем первый пробел после URL
			spaceIndex := -1
			for j := 0; j < len(remainingSlice); j++ {
				if remainingSlice[j] == ' ' {
					spaceIndex = j
					break
				}
			}

			// Определяем конец URL
			urlEndIndex := len(remainingSlice)
			if spaceIndex != -1 {
				urlEndIndex = spaceIndex
			}

			// Создаём слайс звёздочек
			asterisks := make([]byte, urlEndIndex)
			for k := 0; k < urlEndIndex; k++ {
				asterisks[k] = '*'
			}
			// Добавляем звёздочки в результат
			result = append(result, asterisks...)

			// Обновляем текущую позицию
			currentPos = startIndex + urlEndIndex
			// Пропускаем проверку внутри URL
			i = currentPos - 1
		}
	}

	// Добавляем оставшуюся часть строки
	result = append(result, userInputSlice[currentPos:]...)

	return string(result)
}

func main() {
	var userInput string
	fmt.Println("Enter a text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	fmt.Println(receiveUserInput(userInput))
}
