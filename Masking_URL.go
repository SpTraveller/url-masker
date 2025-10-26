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

	// Ищем паттерн в срезе байтов
	patternIndex := -1
	for i := 0; i < len(userInput)-len(patternBytes); i++ {
		match := true
		for j := 0; j < len(patternBytes); j++ {
			if userInputSlice[i+j] != patternBytes[j] {
				match = false
				break
			}
		}
		// Если паттерн не найден
		if match {
			patternIndex = i
			break
		}
	}
	// Начало данных после паттерна
	startIndex := patternIndex + len(patternBytes)
	remainingSlice := userInputSlice[startIndex:]
	// Ищем первый пробел после URL
	spaceIndex := -1
	for i := 0; i < len(remainingSlice); i++ {
		if remainingSlice[i] == ' ' {
			spaceIndex = i
			break
		}
	}
	// Определяем конец URL
	urlEndIndex := len(remainingSlice)
	if spaceIndex != -1 {
		urlEndIndex = spaceIndex
	}
	// Заменяем URL на звездочки
	urlLength := urlEndIndex
	asterisks := ""
	for i := 0; i < urlLength; i++ {
		asterisks += "*"
	}
	// Собираем результат: часть до URL + звездочки + часть с началом пробела
	result := string(userInputSlice[:patternIndex]) + pattern + asterisks
	// Добавляем оставшуюся часть строки (если есть)
	if spaceIndex != -1 {
		result += string(remainingSlice[spaceIndex:])
	}
	return result
}

func main() {
	var userInput string
	fmt.Println("Enter a text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	receiveUserInput(userInput)
	resultOutput := receiveUserInput(userInput)
	fmt.Println("Result:", resultOutput)

}
