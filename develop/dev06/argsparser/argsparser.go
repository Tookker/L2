package argsparser

import "flag"

// Args Структура с входящими аргументами
type Args struct {
	F string
	D string
	S bool
}

// Parse - парсинг входящих аргументов
func Parse() Args {
	var args Args

	flag.StringVar(&args.F, "f", "asd", "выбрать поля, отобразить строки <1,2,3...>, отобразить строки от <2-<число>>, неотображать строки <-2>")
	flag.StringVar(&args.D, "d", "", "использовать другой разделитель")
	flag.BoolVar(&args.S, "s", false, "только строки с разделителем")
	flag.Parse()

	return args
}
