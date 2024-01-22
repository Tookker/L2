package unpack

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	errIncorectStr = errors.New("Incorrect string")
)

type simbolType uint

const (
	digit  simbolType = 0
	escape simbolType = 1
	other  simbolType = 2
)

type combType uint
type combination struct {
	comb combType
	arr  []rune
}

const (
	otherS                   combType = 0 //'a'
	digitS                   combType = 1 //'1'
	otherSAndDigit           combType = 2 //'a1'
	otherSAndEscapeAndEscape combType = 3 //'a\\'
	otherSAndEscapeAndDigit  combType = 4 //'a\1'
	escapeAndEscape          combType = 5 //'\\'
	escapeAndDigit           combType = 6 //'\1'
	digitAndDigit            combType = 6 //'\\1\\5'
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

	return makeString(combArr), nil
}

// getCombinations - получение возможных комбинаций распаковки строк
func getCombinations(str string) ([]combination, error) {
	arr := []rune(str)
	ret := make([]combination, 0, len(arr)+10)

	err := checkFirstSimbol(arr[0])
	if err != nil {
		return nil, errIncorectStr
	}

	for i := 0; i < len(arr); i++ {
		switch getSimbolType(arr[i]) {
		case escape:
			if len(arr) <= i+1 {
				return nil, errIncorectStr
			}

			switch getSimbolType(arr[i+1]) {
			case escape:
				return append(ret, combination{comb: escapeAndEscape, arr: []rune{arr[i]}}), nil
			case digit:
				return append(ret, combination{comb: escapeAndDigit, arr: []rune{arr[i+1]}}), nil
			default:
				return nil, errIncorectStr
			}
		case other:
			if len(arr) <= i+1 {
				return append(ret, combination{comb: otherS, arr: []rune{arr[i]}}), nil
			}
			switch getSimbolType(arr[i+1]) {
			case other:
				ret = append(ret, combination{comb: otherS, arr: []rune{arr[i]}})
			case digit:
				ret = append(ret, combination{comb: otherSAndDigit, arr: []rune{arr[i], arr[i+1]}})
				i++
			case escape:
				if len(arr) <= i+2 {
					return nil, errIncorectStr
				}
				switch getSimbolType(arr[i+2]) {
				case digit:
					if len(arr) <= i+4 {
						return nil, errIncorectStr
					}
					switch getSimbolType(arr[i+4]) {
					case escape:

						ret = append(ret, combination{comb: otherS, arr: []rune{arr[i]}})
						ret = append(ret, combination{comb: escapeAndDigit, arr: []rune{arr[i+2], arr[i+4]}})
					default:
						ret = append(ret, combination{comb: otherS, arr: []rune{arr[i]}})
						ret = append(ret, combination{comb: digitS, arr: []rune{arr[i+2]}})
						i = i + 2
					}
				case escape:
					ret = append(ret, combination{comb: otherSAndEscapeAndEscape, arr: []rune{arr[i]}})
				default:
					return nil, errIncorectStr
				}
			}
		default:
			return nil, errIncorectStr
		}
	}

	return ret, nil
}

func makeString(arr []combination) string {
	resRune := make([]rune, 0, len(arr)+10)

	for i := 0; i < len(arr); i++ {
		switch arr[i].comb {
		case digitAndDigit:
			fallthrough
		case otherSAndDigit:
			val, _ := strconv.Atoi(string(arr[i].arr[1]))
			for j := 0; j < val; j++ {
				resRune = append(resRune, arr[i].arr[0])
			}
		case otherSAndEscapeAndDigit:
			resRune = append(resRune, arr[i].arr[1])
		default:
			resRune = append(resRune, arr[i].arr...)
		}
	}

	return string(resRune)
}

// checkFirstIndx - проверка первого элемента на коректность (не должно быть числа)
func checkFirstSimbol(r rune) error {
	res := getSimbolType(r)
	if res == digit {
		return errIncorectStr
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
