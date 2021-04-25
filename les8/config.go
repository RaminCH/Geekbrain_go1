package main

import (
	"fmt"
	"envconfig"
)

func main() {

	type Config struct {
		DiagPort   int   `port: 8080`
		DB		string	 `db_url: postgres:`//db-user:db-password@petstore-db:5432/petstore?sslmode=disable
		Jurl	string	 `jaeger_url: http:`//jaeger:16686
		Surl	string	 `sentry_url: http:`//sentry:9000
		Kafka	   int	 `kafka_broker: kafka:9092`
		TestID	   int	 `some_app_id: testid`
		TestKey	   int	 `some_app_key: testkey`
	}


	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

}