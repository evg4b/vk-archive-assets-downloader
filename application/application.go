package application

import (
	"log"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/loader"
	"github.com/evg4b/vk-archive-assets-downloader/parser"
)

type Downloader struct {
	wg            *sync.WaitGroup
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	parser        *parser.Parser
	loader        *loader.Loader
}

func NewDownloader(src, dest string, dialogs, types []string) *Downloader {
	dataChanel := make(chan contract.Attachemt)
	wg := sync.WaitGroup{}

	attachemtPb := pb.New(0).Prefix("Attachments")
	dialogsPb := pb.New(0).Prefix("Dialogs")
	dialogPagesPb := pb.New(0).Prefix("Dialog pages")

	log.Printf("Source: %s, Destination: %s", src, dest)
	if dialogs != nil && len(dialogs) > 0 {
		log.Printf("Dialogs: %x", dialogs)
	}
	if types != nil && len(types) > 0 {
		log.Printf("Types: %x", types)
	}
	log.Println()

	return &Downloader{
		wg:            &wg,
		attachemtPb:   attachemtPb,
		dialogsPb:     dialogsPb,
		dialogPagesPb: dialogPagesPb,
		parser: parser.NewParser(&wg, src, dialogs, dataChanel).
			WithAttachemtProgressBar(attachemtPb).
			WithDialogPagesProgressBar(dialogPagesPb).
			WithDialogsProgressBar(dialogsPb),
		loader: &loader.Loader{
			Input: dataChanel,
			Dest:  dest,
			Wg:    &wg,
		},
	}
}
