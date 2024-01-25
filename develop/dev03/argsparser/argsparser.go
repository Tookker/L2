package argsparser

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNotEnoughArgs - мало аргумегтов
	ErrNotEnoughArgs = errors.New("to few arguments")
	// ErrUnknowFlag - неизвестный флаг
	ErrUnknowFlag = errors.New("unknow flag. Allowed flags 'M b c h k n r u'")
	// ErrFlagKIsNegative - флаг К < 0
	ErrFlagKIsNegative = errors.New("flag k cant have a negative digit")
	// ErrAtioErr - ошибка atoi
	ErrAtioErr = errors.New("atoi error")
)

// Flags - структура с возможнами флагами
type Flags struct {
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
}

// ParsedArgs - структура с считыннами аргументами
type ParsedArgs struct {
	flags      Flags
	inputFile  string
	outputFile string
}

// Parse -парсинг входящих аргументов
func Parse(args []string) (ParsedArgs, error) {
	const minArgs = 3
	if len(args) < minArgs {
		return ParsedArgs{}, ErrNotEnoughArgs
	}

	var parsedArgs ParsedArgs
	parsedArgs.inputFile = args[1]
	flags, err := setFlags(args)
	if err != nil {
		return ParsedArgs{}, err
	}

	parsedArgs.flags = flags
	parsedArgs.outputFile = args[len(args)-1]
	return parsedArgs, nil
}

// setFlags - установить флаги
func setFlags(flags []string) (Flags, error) {
	runeFlags := []rune(flags[0])
	var parsedFlags Flags

	for i := 0; i < len(runeFlags); i++ {
		if runeFlags[i] == 'k' {
			if len(runeFlags) <= i+1 {
				res, err := setUintFlag(runeFlags[i], []rune(flags[2])[0], parsedFlags)
				if err != nil {
					return Flags{}, err
				}
				parsedFlags = res
				continue
			} else {
				res, err := setUintFlag(runeFlags[i], runeFlags[i+1], parsedFlags)
				if err != nil {
					return Flags{}, err
				}
				parsedFlags = res
				i++
				continue
			}
		}

		res, err := setBoolFlag(runeFlags[i], parsedFlags)
		if err != nil {
			return Flags{}, err
		}
		parsedFlags = res
	}

	return parsedFlags, nil
}

// setBoolFlag - установить bool akfub
func setBoolFlag(flag rune, flags Flags) (Flags, error) {
	switch flag {
	case 'M':
		flags.M = true
		return flags, nil
	case 'b':
		flags.B = true
		return flags, nil
	case 'c':
		flags.C = true
		return flags, nil
	case 'h':
		flags.H = true
		return flags, nil
	case 'n':
		flags.N = true
		return flags, nil
	case 'r':
		flags.R = true
		return flags, nil
	case 'u':
		flags.U = true
		return flags, nil
	}

	return Flags{}, ErrUnknowFlag
}

// setUintFlag - установить uint флаги
func setUintFlag(flag rune, val rune, flags Flags) (Flags, error) {
	res, err := strconv.Atoi(string(val))
	if err != nil {
		return Flags{}, fmt.Errorf("%w"+err.Error(), ErrAtioErr)
	}

	if res < 0 {
		return Flags{}, ErrFlagKIsNegative
	}

	switch flag {
	case 'k':
		flags.K = res
		return flags, nil
	}

	return Flags{}, ErrUnknowFlag
}
