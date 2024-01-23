package unpack

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	// ErrIncorectStr Ошибка входящей строки
	ErrIncorectStr = errors.New("Incorrect string")
)

type simbolType uint

const (
	digit  simbolType = 0
	escape simbolType = 1
	other  simbolType = 2
)

// String - расскрывает escape последовательность
func String(str string) (string, error) {
	if len(str) == 1 {
		return str, nil
	}

	combArr, err := getCombinations(str)
	if err != nil {
		return "", err
	}

	return combArr, nil
}

// getCombinations - получение возможных комбинаций распаковки строк
func getCombinations(str string) (string, error) {
	arr := []rune(str)
	res := make([]rune, 0, len(str)+10)

	err := checkFirstSimbol(arr[0])
	if err != nil {
		return "", ErrIncorectStr
	}

	for i := 0; i < len(arr); i++ {
		switch getSimbolType(arr[i]) {
		case escape:
			if i+1 >= len(arr) {
				return "", ErrIncorectStr
			}

			if getSimbolType(arr[i+1]) == other {
				return "", ErrIncorectStr
			}

			if getSimbolType(arr[i+1]) == digit {
				if i+2 < len(arr) {
					if getSimbolType(arr[i+2]) == digit {
						iter, resR, err := makeEscapeCombination(arr[i+1 : i+3])
						if err != nil {
							return "", err
						}
						i += iter
						res = append(res, resR...)
						continue
					}
				}

				iter, resR, err := makeEscapeCombination(arr[i+1 : i+2])
				if err != nil {
					return "", err
				}

				i += iter
				res = append(res, resR...)
				continue
			}

			if i+2 >= len(arr) {
				return "", ErrIncorectStr
			}

			if getSimbolType(arr[i+2]) == other {
				iter, resR, err := makeEscapeCombination(arr[i : i+1])
				if err != nil {
					return "", err
				}
				i += iter + 1
				res = append(res, resR...)
				continue
			}

			if getSimbolType(arr[i+2]) == digit {
				iter, resR, err := makeEscapeCombination(arr[i+1 : i+3])
				if err != nil {
					return "", err
				}
				i += iter
				res = append(res, resR...)
				continue
			}

			return "", ErrIncorectStr

		case other:
			if i+1 >= len(arr) {
				_, resR, err := makeOtherCombination(arr[i : i+1])
				if err != nil {
					return "", err
				}
				res = append(res, resR...)
			} else {
				iter, resR, err := makeOtherCombination(arr[i : i+2])
				if err != nil {
					return "", err
				}
				i += iter
				res = append(res, resR...)
			}
		}
	}

	return string(res), nil
}

// makeOtherCombination - создать комбинацию из букв
func makeOtherCombination(arr []rune) (int, []rune, error) {
	const maxSize = 2
	const minSize = 1

	switch len(arr) {
	case maxSize:
		if getSimbolType(arr[1]) == digit {
			size, _ := strconv.Atoi(string(arr[1]))
			res := make([]rune, 0, size)
			for i := 0; i < size; i++ {
				res = append(res, arr[0])
			}
			return 1, res, nil
		}

		return 0, []rune{arr[0]}, nil
	case minSize:
		return 0, []rune{arr[0]}, nil
	}

	return 0, nil, ErrIncorectStr
}

// создать комбинацию из / - escape
func makeEscapeCombination(arr []rune) (int, []rune, error) {
	res := make([]rune, 0, len(arr))

	switch len(arr) {
	case 1:
		res = append(res, arr...)
		return 1, res, nil

	case 2:
		size, _ := strconv.Atoi(string(arr[1]))
		for i := 0; i < size; i++ {
			res = append(res, arr[0])
		}
		return 2, res, nil
	}

	return 0, nil, ErrIncorectStr
}

// checkFirstIndx - проверка первого элемента на коректность (не должно быть числа)
func checkFirstSimbol(r rune) error {
	res := getSimbolType(r)
	if res == digit {
		return ErrIncorectStr
	}

	return nil
}

// getSimbolType - тип символа (буква, цифра, escape-последовательность )
func getSimbolType(r rune) simbolType {
	res := unicode.IsDigit(r)
	if res {
		return digit
	}

	if r == '\\' {
		return escape
	}

	return other
}
