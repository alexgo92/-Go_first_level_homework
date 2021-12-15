package confFile

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"os"
	"strings"
	"v2/internal/config"
)

func ParseConf() {
	flag.Parse()
	filename := flag.Arg(0)

	if _, err := os.Stat(filename); err != nil {
		fmt.Println("file not find")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file can't be open")
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file can't be close")
			return
		}
	}()

	conf := config.Config{}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		fmt.Println("can't be decoded")
	}

	if !govalidator.IsPort(conf.Port) {
		fmt.Println("validation error: wrong port")
		return
	}
	fmt.Println(conf.Port)

	if !govalidator.IsURL(conf.DbUrl) {
		fmt.Println("validation error: wrong DbUrl")
		return
	}
	fmt.Println(conf.DbUrl)

	if !govalidator.IsURL(conf.JaegerUrl) {
		fmt.Println("validation error: wrong JaegerUrl")
		return
	}
	fmt.Println(conf.JaegerUrl)

	if !govalidator.IsURL(conf.SentryUrl) {
		fmt.Println("validation error: wrong SentryUrl")
		return
	}
	fmt.Println(conf.SentryUrl)

	// валидация KafkaBroker
	words := strings.Split(conf.KafkaBroker, ":")
	if len(words) != 2 {
		fmt.Println("validation error: wrong KafkaBroker")
		return
	}
	firstWord := words[0]
	secondWord := words[1]
	if firstWord != "kafka" || !govalidator.IsPort(secondWord) {
		fmt.Println("validation error: wrong KafkaBroker")
		return
	}
	fmt.Println(conf.KafkaBroker)

	fmt.Println(conf.SomeAppId)
	fmt.Println(conf.SomeAppKey)
}
