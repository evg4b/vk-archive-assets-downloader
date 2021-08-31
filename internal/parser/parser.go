package parser

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/internal/common"
)

const dir = "messages"

type Parser struct {
	path          string
	ids           []string
	output        chan<- common.Attachemt
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	wg            *sync.WaitGroup
}

func NewParser(wg *sync.WaitGroup, path string, ids []string, output chan<- common.Attachemt) *Parser {
	return &Parser{
		path:   path,
		ids:    ids,
		output: output,
		wg:     wg,
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

func (p *Parser) Parse(ctx context.Context) {
	defer p.wg.Done()
	defer close(p.output)

	dirs, err := p.load()
	if err != nil {
		panic(err)
	}

	log.Println("Parser started")

	for _, dirPath := range dirs {
		files, err := p.parseDialog(dirPath)
		if err != nil {
			log.Printf("ERROR: failed to read dis %s\n", dirPath)
			continue
		}

		common.InitProgressBar(p.dialogPagesPb, len(files))
		dialogName, err := p.getDialogName(files[0])
		if err != nil {
			log.Printf("ERROR: failed to read dis %s\n", dirPath)
			continue
		}

		for _, filePath := range files {
			p.processFile(dialogName, filePath)
		}
	}
}

func (p *Parser) load() ([]string, error) {
	folderPath := path.Join(p.path, dir)
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, folder := range folders {
		if folder.IsDir() && common.IncludeOrEmpty(folder.Name(), p.ids) {
			paths = append(paths, filepath.Join(folderPath, folder.Name()))
		}
	}

	common.InitProgressBar(p.dialogsPb, len(paths))

	return paths, nil
}
