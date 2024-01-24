package anagram

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

var (
	// ErrLetterIsNotRussian - не русское слово
	ErrLetterIsNotRussian = errors.New("allow only russian letters")
	// ErrDigitInWord - в слове есть цифра
	ErrDigitInWord = errors.New("digit in word")
)

// Make - создать мапу анаграмм
func Make(words []string) (map[string][]string, error) {
	//Проверяем все слова, что они на русском языке
	var err error
	for _, word := range words {
		err = isRussianWord(word)
		if err != nil {
			return nil, err
		}
	}

	tempMap := make(map[string][]string, len(words))
	var lowWord string

	for _, word := range words {
		lowWord = strings.ToLower(word)

		_, ok := tempMap[lowWord]
		if ok {
			continue
		}

		key, ok := findAnagramKey(lowWord, tempMap)
		if ok {
			if containsWord(lowWord, tempMap[key]) {
				continue
			}
			tempMap[key] = append(tempMap[key], lowWord)
			sort.Strings(tempMap[key])
			continue
		}

		tempMap[lowWord] = []string{}
	}

	return makeRes(tempMap), nil
}

// isRussian - проверка что слово на русском языке
func isRussianWord(word string) error {
	for _, letter := range word {
		if unicode.IsDigit(rune(letter)) {
			return ErrDigitInWord
		}

		if !unicode.Is(unicode.Cyrillic, rune(letter)) {
			return ErrLetterIsNotRussian
		}
	}

	return nil
}

// getAnagram - получить анаграмму слова
func getAnagram(word string) string {
	rWord := []rune(word)
	sort.Slice(rWord, func(i, j int) bool { return rWord[i] < rWord[j] })
	return string(rWord)
}

// findAnagramKey - поиск анаграммы слова в мапе
func findAnagramKey(word string, anagramMap map[string][]string) (string, bool) {
	anagramWord := getAnagram(word)
	var anagramKey string

	for key := range anagramMap {
		anagramKey = getAnagram(key)
		if anagramKey == anagramWord {
			return key, true
		}
	}

	return "", false
}

// makeRes - сформировать результат
func makeRes(anagramMap map[string][]string) map[string][]string {
	res := make(map[string][]string)

	for key, val := range anagramMap {
		if len(val) == 0 {
			continue
		}

		res[key] = append(res[key], val...)
	}

	return res
}

// containsWord - проверка на повторение слов
func containsWord(word string, arr []string) bool {
	for _, val := range arr {
		if val == word {
			return true
		}
	}

	return false
}
