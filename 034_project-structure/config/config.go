package config

import (
	"fmt"
	"os"
	"strconv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func loadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}

	ServiceName := os.Getenv("SERVICE_NAME")
	if ServiceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPOrt := os.Getenv("HTTP_PORT")
	if httpPOrt == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPOrt, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: ServiceName,
		HttpPort:    port,
	}

}

func GetConfig() Config{
	loadConfig()
	return configuration
}