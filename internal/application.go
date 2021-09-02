package internal

import (
	"context"
	"log"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/internal/common"
	"github.com/evg4b/vk-archive-assets-downloader/internal/loader"
	"github.com/evg4b/vk-archive-assets-downloader/internal/parser"
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
	dataChanel := make(chan common.Attachemt)
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

func (d *Downloader) Run(ctx context.Context) error {
	pool, err := pb.StartPool(d.dialogsPb, d.dialogPagesPb, d.attachemtPb)
	if err != nil {
		return err
	}

	defer pool.Stop()

	d.wg.Add(2)
	go d.parser.Parse(ctx)
	go d.loader.Load(ctx)

	d.wg.Wait()

	return nil
}
