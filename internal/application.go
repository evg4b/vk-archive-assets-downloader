package internal

import (
	"sync"

	"github.com/cheggaaa/pb"
)

type Downloader struct {
	src           string
	dialogs       []string
	types         []string
	dest          string
	wg            sync.WaitGroup
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
}

func NewDownloader(src, dest string, dialogs, types []string) *Downloader {
	return &Downloader{
		src:           src,
		dest:          dest,
		dialogs:       dialogs,
		types:         types,
		wg:            sync.WaitGroup{},
		attachemtPb:   pb.New(0).Prefix("Attachments"),
		dialogsPb:     pb.New(0).Prefix("Dialogs"),
		dialogPagesPb: pb.New(0).Prefix("Dialog pages"),
	}
}

func (d *Downloader) Run() error {
	pool, err := pb.StartPool(d.dialogsPb, d.dialogPagesPb, d.attachemtPb)
	if err != nil {
		panic(err)
	}
	defer pool.Stop()

	println(d.src)
	println(d.dialogs)
	println(d.types)
	println(d.dest)

	d.wg.Wait()

	return nil
}
