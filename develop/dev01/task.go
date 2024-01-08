package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
1.	Базовая задача

Создать программу печатающую точное время с использованием NTP -библиотеки.
Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
1.	Программа должна быть оформлена как go module
2.	Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS
*/

// GetTime - функция получения текущего времени, используя github.com/beevik/ntp библиотеку
func GetTime(addres string) (time.Time, error) {
	return ntp.Time(addres)
}

func main() {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(time)
}
