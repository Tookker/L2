package cut

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"L2/develop/dev06/argsparser"
)

// MyCut - структура MyCut
type MyCut struct {
	skipStr int
	showStr []uint
	from    int
	to      int
	d       string
	s       bool
	str     string
}

var (
	// ErrParseArgF - ошибка парсинга флага f
	ErrParseArgF = errors.New("error parse flag f")
	//ErrParseArgFNegative - ошибка парсинга флага f - числа не могут быть отрицательными
	ErrParseArgFNegative = errors.New("f flag cant be negative")
)

// NewCut - конструктор структуры MyCut
func NewCut(args argsparser.Args, str string) (MyCut, error) {
	cut := MyCut{
		s:   args.S,
		str: str,
	}

	if args.D == "" {
		cut.d = "\t"
	} else {
		cut.d = args.D
	}

	err := cut.readArgF(args.F)
	if err != nil {
		return MyCut{}, err
	}

	return cut, nil
}

// Cut - выполнить сut комманду
func (m *MyCut) Cut() string {
	res := m.getWorkColumn()
	res = m.delemite(res)
	return m.getRes(res)
}

// getWorkColumn - убрать неиспользуемые колонки
func (m *MyCut) getWorkColumn() [][]string {
	splitEnter := strings.Split(m.str, "\n")
	res := make([][]string, 0, cap(splitEnter))
	for _, val := range splitEnter {
		splitSpace := strings.Split(val, " ")
		res = append(res, splitSpace)
	}

	if len(m.showStr) != 0 {
		tmp := make([][]string, len(splitEnter))
		for i := range res {
			tmp[i] = make([]string, 0, cap(res[i]))
		}
		copy(tmp, res)
		clear(res)

		for indx, node := range tmp {
			for i, val := range node {
				if !slices.Contains(m.showStr, uint(i+1)) {
					continue
				}

				res[indx] = append(res[indx], val)
			}
		}
	} else if m.skipStr != 0 {
		tmp := make([][]string, len(splitEnter))
		for i := range res {
			tmp[i] = make([]string, 0, cap(res[i]))
		}
		copy(tmp, res)
		clear(res)

		for indx, node := range tmp {
			for i, val := range node {
				if i+1 == m.skipStr {
					continue
				}

				res[indx] = append(res[indx], val)
			}
		}
	} else if m.from != 0 && m.to != 0 {
		tmp := make([][]string, len(splitEnter))
		for i := range res {
			tmp[i] = make([]string, 0, cap(res[i]))
		}
		copy(tmp, res)
		clear(res)

		for indx, node := range tmp {
			for i, val := range node {
				if i+1 < m.from || i+1 > m.to {
					continue
				}

				res[indx] = append(res[indx], val)
			}
		}
	}

	return res
}

// delemite - разделить входящую строку на подстроки и вывести в мапе
func (m *MyCut) delemite(sl [][]string) [][]string {
	res := make([][]string, len(sl))

	for index, node := range sl {
		res[index] = make([]string, 0, len(node))
		for _, str := range node {
			if m.s {
				if !strings.Contains(str, m.d) {
					continue
				}
			}
			res[index] = append(res[index], strings.ReplaceAll(str, m.d, ""))
		}
	}
	return res
}

// readArgF - чтенине флага f
func (m *MyCut) readArgF(args string) error {
	if len(args) == 0 {
		return nil
	} else if len(args) == 1 {
		num, err := strconv.Atoi(args)
		if err != nil {
			return fmt.Errorf("%w vals "+args+": "+err.Error(), ErrParseArgF)
		}
		m.showStr = append(m.showStr, uint(num))
		return nil
	}

	if strings.Contains(args, ",") {
		res := strings.Split(args, ",")
		for _, val := range res {
			num, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("%w vals "+args+": "+err.Error(), ErrParseArgF)
			}

			if num < 0 {
				return fmt.Errorf("%w val"+val, ErrParseArgFNegative)
			}

			m.showStr = append(m.showStr, uint(num))
		}
		return nil
	}

	runeArr := []rune(args)
	if runeArr[0] == '-' {
		arr := string(runeArr[1:])
		num, err := strconv.Atoi(arr)
		if err != nil {
			return fmt.Errorf("%w vals "+arr+": "+err.Error(), ErrParseArgF)
		}

		if num < 0 {
			return ErrParseArgFNegative
		}

		m.skipStr = num
		return nil
	} else if strings.Contains(args, "-") {
		res := strings.Split(args, "-")
		if len(res) != 2 {
			return ErrParseArgF
		}

		for i, val := range res {
			num, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("%w val "+val+": "+err.Error(), ErrParseArgF)
			}

			if num < 0 {
				return ErrParseArgFNegative
			}

			if i == 0 {
				m.from = num
			} else if i == 1 {
				m.to = num
			}
		}

		return nil
	}

	return ErrParseArgF
}

// getRes - вывод результата работы cut
func (m *MyCut) getRes(strs [][]string) string {
	var res []rune
	for _, str := range strs {
		res = append(res, []rune(strings.Join(str, " "))...)
		res = append(res, '\n')
	}

	return string(res)
}
