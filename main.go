package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/evg4b/vk-archive-assets-downloader/application"
)

func main() {
	logfile := openLogFile()
	defer logfile.Close()

	log.SetOutput(logfile)

	app := application.NewDownloader()
	app.ParseArguments()
	err := app.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func openLogFile() *os.File {
	logfile := fmt.Sprintf("vk-archive-assets-downloader-%s.log", time.Now().Format("20060102150405"))
	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return file
}
