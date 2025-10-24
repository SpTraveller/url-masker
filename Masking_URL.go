package main

import (
	"bufio"
	"fmt"
	"os"
)

// Структура для хранения информации о найденном URL:
type URLInfo struct {
	Protocol   string // "http://" или "https://"
	URL        []byte // URL без протокола
	StartIndex int    // позиция начала протокола в исходном тексте
	EndIndex   int    // позиция конца URL в исходном тексте
}

// функция для принятия ввода и форматирования его в байты
func getUserInput() []byte {
	fmt.Println("Enter a text:")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	//удаление переноса строки
	if len(userInput) > 0 && userInput[len(userInput)-1] == '\n' {
		userInput = userInput[:len(userInput)-1]
	}
	return []byte(userInput)
}

// считывание префикса:
func checkPrefix(address []byte, prefix string) bool {
	prefixBytes := []byte(prefix)
	if len(address) < len(prefixBytes) {
		return false
	}
	for i := range prefixBytes {
		if address[i] != prefixBytes[i] {
			return false
		}
	}
	return true
}

// проверка символов
func checkUrlSymbols(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9') ||
		b == '-' || b == '_' ||
		b == '/' || b == '?' || b == '#' ||
		b == '@' || b == ':' || b == '='
}

// проверка символа - буква или цифра
func checkCharOrNum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

// Находим и вычленяем URL из текста
func findUrl(text []byte) *URLInfo {
	// Проверка на пустой ввод
	if len(text) == 0 {
		return nil
	}

	for i := 0; i <= len(text)-len("http://"); i++ {
		// если протокол в начале - https://
		if checkPrefix(text[i:], "https://") {
			protocol := "https://"
			endIndex := i + len(protocol)
			// Читаем URL с обработкой точки
			for endIndex < len(text) {
				currentChar := text[endIndex]

				// Если это базовый URL-символ — включаем
				if checkUrlSymbols(currentChar) {
					endIndex++
					continue
				}

				// Если точка — проверяем, что после неё
				if currentChar == '.' {
					// Проверяем, есть ли символ после точки
					if endIndex+1 < len(text) {
						nextChar := text[endIndex+1]
						// Если после точки буква или цифра — это часть URL
						if checkCharOrNum(nextChar) {
							endIndex++
							continue
						}
					}
					// Иначе точка в конце — останавливаемся
					break
				}
				// Любой другой символ — конец URL
				break
			}
			extractedUrl := text[i+len(protocol) : endIndex]
			return &URLInfo{protocol, extractedUrl, i, endIndex}
		}

		// если протокол в начале - http://
		if checkPrefix(text[i:], "http://") {
			protocol := "http://"
			endIndex := i + len(protocol)

			for endIndex < len(text) {
				currentChar := text[endIndex]

				if checkUrlSymbols(currentChar) {
					endIndex++
					continue
				}

				if currentChar == '.' {
					if endIndex+1 < len(text) {
						nextChar := text[endIndex+1]
						if checkCharOrNum(nextChar) {
							endIndex++
							continue
						}
					}
					break
				}
				break
			}
			extractedUrl := text[i+len(protocol) : endIndex]
			return &URLInfo{protocol, extractedUrl, i, endIndex}
		}
	}
	return nil
}

// функция создания слайса маски
func createMask(length int) []byte {
	mask := make([]byte, length)
	for i := 0; i < length; i++ {
		mask[i] = '*'
	}
	return mask
}

// функция замены URL на маску
func maskUrlInText(text []byte, urlInfo *URLInfo) []byte {
	// Создаём маску из звёздочек длиной исходного URL
	mask := createMask(len(urlInfo.URL))

	// Вырезаем часть до URL (включая протокол)
	cutoutLeft := text[:urlInfo.StartIndex+len([]byte(urlInfo.Protocol))]

	// Вырезаем часть после URL
	cutoutRight := text[urlInfo.EndIndex:]

	// Склейка: часть текста до URL + маска + часть после
	output := make([]byte, 0, len(cutoutLeft)+len(mask)+len(cutoutRight))
	output = append(output, cutoutLeft...)
	output = append(output, mask...)
	output = append(output, cutoutRight...)

	return output

}

func main() {

	inputText := getUserInput()
	fmt.Println(string(inputText))

	// Проверка на пустой ввод
	if len(inputText) == 0 {
		fmt.Println("Empty input")
		return
	}

	urlInfo := findUrl(inputText)

	// Проверка на присутствие URL в тексте
	if urlInfo == nil {
		fmt.Println("URL not found text")
		return
	}
	fmt.Println(string(urlInfo.URL))

	maskedURL := maskUrlInText(inputText, urlInfo)
	fmt.Println(string(maskedURL))

}
