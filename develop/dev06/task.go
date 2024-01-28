package main

import (
	"bufio"
	"fmt"
	"os"

	"L2/develop/dev06/argsparser"
	"L2/develop/dev06/cut"
)

/*
6.	Утилита cut

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

func main() {
	args := argsparser.Parse()
	fmt.Println(args)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	cut, err := cut.NewCut(args, input.Text())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(cut.Cut())
}
