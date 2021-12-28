package main

import (
	"etl/logger"
	"etl/translator"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Logger *logger.LogComponent
)

func main() {
	Logger = logger.Logger

	err := godotenv.Load()
	if err != nil {
		Logger.LogError(err.Error())
		os.Exit(1)
	}

	args := os.Args[1:]
	for _, arg := range args {
		executeFetch(arg)
	}
}

func executeFetch(endpoint string) {
	switch endpoint {
	case "parking":
		translator.FetchParking()
	case "station":
		translator.FetchStation()
	case "traffic":
		translator.FetchTraffic()
	case "weather":
		translator.FetchWeather()
	default:
		fmt.Printf("Invalid Argument: %s\n", endpoint)
	}
}
