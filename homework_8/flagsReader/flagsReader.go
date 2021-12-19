package flagsReader

import (
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"strings"
)

var (
	Port        = flag.String("port", "", "this is the port to use")
	DbUrl       = flag.String("db_url", "", "path to database")
	JaegerUrl   = flag.String("jaeger_url", "", "address jaeger_url")
	SentryUrl   = flag.String("sentry_url", "", "address sentry_url")
	KafkaBroker = flag.String("kafka_broker", "", "port kafka_broker")
	SomeAppId   = flag.String("some_app_id", "", "something_id")
	SomeAppKey  = flag.String("some_app_key", "", "something_key")
)

func ParseConf() {
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic: ", err)
		}
	}()

	if !govalidator.IsPort(*Port) {
		fmt.Println("validation error: wrong port")
		return
	}
	fmt.Println(*Port)

	if !govalidator.IsURL(*DbUrl) {
		fmt.Println("validation error: wrong DbUrl")
		return
	}
	fmt.Println(*DbUrl)

	if !govalidator.IsURL(*JaegerUrl) {
		fmt.Println("validation error: wrong JaegerUrl")
		return
	}
	fmt.Println(*JaegerUrl)

	if !govalidator.IsURL(*SentryUrl) {
		fmt.Println("validation error: wrong SentryUrl")
		return
	}
	fmt.Println(*SentryUrl)

	// валидация KafkaBroker
	words := strings.Split(*KafkaBroker, ":")
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
	fmt.Println(*KafkaBroker)

	// тут валидация не нужна вроде как
	fmt.Println(*SomeAppId)
	fmt.Println(*SomeAppKey)
}

// Команда для проверки
/* go run main.go --scanMethod=flag --port=8082 --db_url=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaeger_url=http://jaeger:16686 --sentry_url=http://sentry:9000 --kafka_broker=kafka:9092 --some_app_id=testid --some_app_key=testkey  */
