package config

type Config struct {
	Port        string `json:"port" envconfig:"PORT" default:"8080" required:"true"`
	DbUrl       string `json:"dbUrl" envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerUrl   string `json:"jaegerUrl" envconfig:"JAEGER_URL" default:"http://jaeger:16686" required:"true"`
	SentryUrl   string `json:"sentryUrl" envconfig:"SENTRY_URL" default:"http://sentry:9000" required:"true"`
	KafkaBroker string `json:"kafkaBroker" envconfig:"KAFKA_BROKER" default:"kafka:9092" required:"true"`
	SomeAppId   string `json:"someAppId" envconfig:"SOME_APP_ID" default:"testid" required:"true"`
	SomeAppKey  string `json:"someAppKey" envconfig:"SOME_APP_KEY" default:"testkey" required:"true"`
}
