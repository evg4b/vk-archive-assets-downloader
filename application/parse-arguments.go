package application

import (
	"flag"
	"log"
	"runtime"

	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/loader"
	"github.com/evg4b/vk-archive-assets-downloader/parser"
)

func (d *Downloader) ParseArguments() {
	src := flag.String("src", "archive", "path to archive folder")
	dialogs := flag.String("dialogs", "", "coma separeted dialogs ids")
	types := flag.String("types", "", "coma separeted attachments types")
	dest := flag.String("dest", "dest", "destination folder")
	encoding := flag.String("encoding", "Windows1251", "destination folder")
	poolSize := flag.Int("pool", 1000, "attachments pool size")
	threadsCount := flag.Int("threads", runtime.NumCPU(), "Loader threads count")

	flag.Parse()

	log.Printf("Source: %s, Destination: %s", *src, *dest)
	if dialogs != nil && len(*dialogs) > 0 {
		log.Printf("Dialogs: %s", *dialogs)
	}

	if types != nil && len(*types) > 0 {
		log.Printf("Types: %s", *types)
	}

	log.Printf("Encoding: %s", *encoding)
	log.Printf("Pool size: %d", *poolSize)
	log.Printf("Threads count: %d", *threadsCount)

	dataChanel := make(chan contract.Attachemt, *poolSize)

	d.parser = parser.NewParser(
		dataChanel,
		parser.Source(*src),
		parser.Encoding(*encoding),
		parser.Dialogs(*dialogs),
		parser.Types(*types),
		parser.AttachemtProgressBar(d.attachemtPb),
		parser.DialogPagesProgressBar(d.dialogPagesPb),
		parser.DialogsProgressBar(d.dialogsPb),
	)

	d.loader = loader.NewLoader(
		dataChanel,
		loader.Destination(*dest),
		loader.AttachemtProgressBar(d.attachemtPb),
		loader.ThreadsCount(*threadsCount),
	)
}
