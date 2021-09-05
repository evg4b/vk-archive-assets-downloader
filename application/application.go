package application

import (
	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/loader"
	"github.com/evg4b/vk-archive-assets-downloader/parser"
)

type Downloader struct {
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	parser        *parser.Parser
	loader        *loader.Loader
}

func NewDownloader() *Downloader {
	return &Downloader{
		attachemtPb:   pb.New(0).Prefix("Attachments"),
		dialogsPb:     pb.New(0).Prefix("Dialogs"),
		dialogPagesPb: pb.New(0).Prefix("Dialog pages"),
	}
}
