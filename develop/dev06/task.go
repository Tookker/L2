package main

import (
	"bufio"
	"fmt"
	"os"

	"L2/develop/dev06/argsparser"
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

	fmt.Println([]rune(input.Text()))
}
