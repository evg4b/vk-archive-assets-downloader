package parser

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils/collections"
	"github.com/evg4b/vk-archive-assets-downloader/utils/progressbar"
)

const dir = "messages"

type Parser struct {
	path          string
	ids           []string
	output        chan<- contract.Attachemt
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	wg            *sync.WaitGroup
	log           *log.Logger
}

func NewParser(wg *sync.WaitGroup, path string, ids []string, output chan<- contract.Attachemt) *Parser {
	return &Parser{
		path:   path,
		ids:    ids,
		output: output,
		wg:     wg,
		log:    log.New(log.Writer(), "Parser |", log.Flags()),
	}
}

func (p *Parser) WithAttachemtProgressBar(progressBar *pb.ProgressBar) *Parser {
	p.attachemtPb = progressBar

	return p
}

func (p *Parser) WithDialogsProgressBar(progressBar *pb.ProgressBar) *Parser {
	p.dialogsPb = progressBar

	return p
}

func (p *Parser) WithDialogPagesProgressBar(progressBar *pb.ProgressBar) *Parser {
	p.dialogPagesPb = progressBar

	return p
}

func (p *Parser) load() ([]string, error) {
	folderPath := path.Join(p.path, dir)
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, folder := range folders {
		if folder.IsDir() && collections.IncludeOrEmpty(folder.Name(), p.ids) {
			paths = append(paths, filepath.Join(folderPath, folder.Name()))
		}
	}

	progressbar.InitProgressBar(p.dialogsPb, len(paths))

	return paths, nil
}
