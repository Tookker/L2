package main

import (
	"fmt"
	"os"

	"L2/develop/dev02/unpack"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
	●	"a4bc2d5e" => "aaaabccddddde"
	●	"abcd" => "abcd"
	●	"45" => "" (некорректная строка)
	●	"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
	●	qwe\4\5 => qwe45 (*)
	●	qwe\45 => qwe44444 (*)
	●	qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	res, err := unpack.String(args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Input string:", args[1], "Output string:", res)
}
