package main

import (
	"flag"
	"strings"

	"github.com/evg4b/vk-archive-assets-downloader/internal"
)

func main() {
	src := flag.String("src", "archive", "path to archive folder")
	dialogs := flag.String("dialogs", "", "coma separeted dialogs ids")
	types := flag.String("types", "", "coma separeted attachments types")
	dest := flag.String("dest", "dest", "destination folder")

	flag.Parse()

	app := internal.NewDownloader(*src, *dest, strings.Split(*dialogs, ","), strings.Split(*types, ","))
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
