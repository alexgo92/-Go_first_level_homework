package main

import (
	"flag"
	"fmt"
	"v2/confFile"
	"v2/envReader"
	"v2/flagsReader"
)

var ScanMethod = flag.String("scanMethod", "", "Указываем откуда хотим считать переменные: "+
	"через флаги или переменные окружения или через файл (обязательно указываем имя файла)")

func main() {
	flag.Parse()

	//Выбор как считываем конфигурацию: либо через флаги, либо через переменные окружения
	if *ScanMethod == "flag" {
		flagsReader.ParseConf()
	} else if *ScanMethod == "env" {
		envReader.ParseConf()
	} else if *ScanMethod == "file" {
		confFile.ParseConf()
	} else {
		fmt.Println("Вы ввели неверное значение флага scanMethod")
	}
}
