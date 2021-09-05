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
)

const dir = "messages"

type ParserOption = func(parser *Parser)

type Parser struct {
	path          string
	encoding      string
	ids           []string
	types         []string
	output        chan<- contract.Attachemt
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	wg            *sync.WaitGroup
	log           *log.Logger
}

func NewParser(output chan<- contract.Attachemt, options ...ParserOption) *Parser {
	parser := &Parser{
		path:     "src",
		encoding: "Windows1251",
		ids:      []string{},
		output:   output,
		wg:       &sync.WaitGroup{},
		log:      log.New(log.Writer(), "Parser |", log.Flags()),
	}

	if options != nil {
		for _, option := range options {
			option(parser)
		}
	}

	return parser
}

func (p *Parser) Wait() {
	p.wg.Wait()
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

	p.dialogsPb.Finish()
	p.dialogsPb.Reset(len(paths))

	return paths, nil
}
