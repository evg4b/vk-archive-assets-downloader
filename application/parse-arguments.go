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
	albums := flag.String("albums", "", "coma separeted albums ids")
	types := flag.String("types", "", "coma separeted attachments types")
	dest := flag.String("dest", "dest", "destination folder")
	encoding := flag.String("encoding", "Windows1251", "destination folder")
	poolSize := flag.Int("pool", 1000, "attachments pool size")
	threadsCount := flag.Int("threads", runtime.NumCPU(), "Loader threads count")

	flag.Parse()

	log.Printf("source: %s, destination: %s", *src, *dest)
	if dialogs != nil && len(*dialogs) > 0 {
		log.Printf("dialogs: %s", *dialogs)
	}

	if albums != nil && len(*albums) > 0 {
		log.Printf("albums: %s", *albums)
	}

	if types != nil && len(*types) > 0 {
		log.Printf("types: %s", *types)
	}

	log.Printf("encoding: %s", *encoding)
	log.Printf("pool size: %d", *poolSize)
	log.Printf("threads count: %d", *threadsCount)

	dataChanel := make(chan contract.Attachemt, *poolSize)

	d.parser = parser.NewParser(
		dataChanel,
		parser.Source(*src),
		parser.Encoding(*encoding),
		parser.Dialogs(*dialogs),
		parser.Albums(*albums),
		parser.Types(*types),
		parser.AttachemtProgressBar(d.attachemtPb),
		parser.DialogPagesProgressBar(d.dialogPagesPb),
		parser.DialogsProgressBar(d.dialogsPb),
		parser.AlbumsProgressBar(d.albumsPb),
	)

	d.loader = loader.NewLoader(
		dataChanel,
		loader.Destination(*dest),
		loader.AttachemtProgressBar(d.attachemtPb),
		loader.ThreadsCount(*threadsCount),
	)
}
