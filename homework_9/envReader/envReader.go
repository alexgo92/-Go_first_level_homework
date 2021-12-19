package envReader

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/kelseyhightower/envconfig"
	"strings"
	"v2/internal/config"
)

func ParseConf() {
	conf := config.Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		fmt.Println(err.Error())
		return
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

	// тут валидация не нужна вроде как
	fmt.Println(conf.SomeAppId)
	fmt.Println(conf.SomeAppKey)
}
