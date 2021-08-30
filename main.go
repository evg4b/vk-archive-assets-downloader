package main

import (
	"context"
	"flag"

	"github.com/evg4b/vk-archive-assets-downloader/internal"
	"github.com/evg4b/vk-archive-assets-downloader/internal/common"
)

func main() {
	src := flag.String("src", "archive", "path to archive folder")
	dialogs := flag.String("dialogs", "", "coma separeted dialogs ids")
	types := flag.String("types", "", "coma separeted attachments types")
	dest := flag.String("dest", "dest", "destination folder")

	flag.Parse()

	app := internal.NewDownloader(*src, *dest, common.SplitNotEmpty(*dialogs), common.SplitNotEmpty(*types))
	err := app.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
