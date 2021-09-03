package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/evg4b/vk-archive-assets-downloader/application"
	"github.com/evg4b/vk-archive-assets-downloader/utils/collections"
)

func main() {
	src := flag.String("src", "archive", "path to archive folder")
	dialogs := flag.String("dialogs", "", "coma separeted dialogs ids")
	types := flag.String("types", "", "coma separeted attachments types")
	dest := flag.String("dest", "dest", "destination folder")

	flag.Parse()

	logfile := openLogFile()
	defer logfile.Close()

	log.SetOutput(logfile)

	app := application.NewDownloader(*src, *dest, collections.SplitNotEmpty(*dialogs), collections.SplitNotEmpty(*types))
	err := app.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func openLogFile() *os.File {
	logfile := fmt.Sprintf("%s.log", time.Now().Format("20060102150405"))
	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return file
}
