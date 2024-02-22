package main

import (
	"fmt"
	"time"
)

/*
7.	Or channel

Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.
*/

func OrChan(channels ...<-chan interface{}) <-chan interface{} {
	mergeChan := make(chan interface{})

	go func() {
		defer close(mergeChan)

		for i := range channels {
			select {
			case <-channels[i]:
				mergeChan <- channels[i]
			default:
				continue
			}
		}
	}()

	return mergeChan
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-OrChan(
		sig(5*time.Second),
		sig(4*time.Second),
		sig(3*time.Second),
		sig(2*time.Second),
		sig(1*time.Second),
	)

	fmt.Println("fone after", time.Since(start))
}
