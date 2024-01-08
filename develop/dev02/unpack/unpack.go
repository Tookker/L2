package unpack

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	incorectStrErr = errors.New("Incorrect string")
)

type simbolType uint

const (
	letter    simbolType = 0
	digit     simbolType = 1
	escape    simbolType = 2
	undefined simbolType = 3
)

func String(str string) (string, error) {
	runeStr := []rune(str)

	err := checkSimbols(runeStr)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(runeStr); i++ {
		if i == 0 {
			err := checkFirstSimbol(runeStr[i])
			return "", err
		}

		if i+1 >= len(runeStr) {
			return string(append(runeStr, runeStr[i])), nil
		}

		switch getSimbolType(runeStr[i]) {
		case letter:
			switch getSimbolType(runeStr[i+1]) {
			case letter:
				runeStr = append(runeStr, runeStr[i])
			case digit:
				val, _ := strconv.Atoi(string(runeStr[i+1]))
				tempArr := make([]rune, 0, val)
				for j := 0; j < val; j++ {
					tempArr = append(tempArr, runeStr[i])
				}
				runeStr = append(runeStr, tempArr...)
				i++
			case escape:
				if i+1 >= len(runeStr) {
					return string(append(runeStr, runeStr[i])), nil
				}
				switch getSimbolType(runeStr[i+2]) {
				case digit:
					runeStr = append(runeStr, runeStr[i], runeStr[i+2])
					i = i + 2
				case escape:
					runeStr = append(runeStr, runeStr[i])
					i = i + 2
				default:
					return "", incorectStrErr
				}
			}
		case escape:
			switch getSimbolType(runeStr[i+1]) {
			case digit:
				runeStr = append(runeStr, runeStr[i+1])
				i++
			case escape:
				runeStr = append(runeStr, runeStr[i])
				i = i + 2
			default:
				return "", incorectStrErr
			}
		default:
			return "", incorectStrErr
		}

		switch getSimbolType(runeStr[i+1]) {
		case letter:
			runeStr = append(runeStr, runeStr[i])
		case digit:
			val, _ := strconv.Atoi(string(runeStr[i+1]))
			tempArr := make([]rune, 0, val)
			for j := 0; j < val; j++ {
				tempArr = append(tempArr, runeStr[i])
			}
			runeStr = append(runeStr, tempArr...)
			i++
		case escape:

		}

	}

	return "", nil
}

// checkFirstIndx - проверка первого элемента на коректность (не должно быть числа)
func checkFirstSimbol(r rune) error {
	res := unicode.IsDigit(r)
	if res {
		return incorectStrErr
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

	res = unicode.IsLetter(r)
	if res {
		return letter
	}

	return undefined
}

func checkSimbols(r []rune) error {
	for _, val := range r {
		res := getSimbolType(val)
		if res == undefined {
			return incorectStrErr
		}
	}

	return nil
}
