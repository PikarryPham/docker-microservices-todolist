package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/google/uuid"
	"golangbook/config"
	"golangbook/consul"

	"golangbook/router"
)

var myUUID = uuid.New()

//go:embed .env
var envVarsFile embed.FS

func main() {
	config.EnvVarsFile = envVarsFile
	r := router.Router()

	_, err := config.AppConfig()
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("\nShutting down.")
		os.Exit(0)
	}()

	consul.ServiceRegistryWithConsul(config.IPAddress, config.ServerPort, myUUID)

	fmt.Printf("Starting Hello Server: %v:%v", config.IPAddress, config.ServerPort)

	http.ListenAndServe(":"+strconv.Itoa(config.ServerPort), r)
}
