package parser

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils"
	"golang.org/x/sync/errgroup"
)

const dir = "messages"
const photos_dir = "photos"
const photo_albums_dir = "photo-albums"

type ParserOption = func(parser *Parser)

type Parser struct {
	path          string
	encoding      string
	ids           []string
	albumsIds     []string
	types         []string
	output        chan<- contract.Attachemt
	attachemtPb   *pb.ProgressBar
	dialogsPb     *pb.ProgressBar
	albumsPb      *pb.ProgressBar
	dialogPagesPb *pb.ProgressBar
	logger        *log.Logger
	errGroup      *errgroup.Group
}

func NewParser(output chan<- contract.Attachemt, options ...ParserOption) *Parser {
	parser := &Parser{
		path:      "src",
		encoding:  "Windows1251",
		ids:       []string{},
		albumsIds: []string{},
		output:    output,
		logger:    log.New(log.Writer(), "[parser] ", log.Flags()),
		errGroup:  &errgroup.Group{},
	}

	for _, option := range options {
		option(parser)
	}

	return parser
}

func (p *Parser) Wait() error {
	return p.errGroup.Wait()
}

func (p *Parser) loadDialogs() ([]string, error) {
	folderPath := path.Join(p.path, dir)
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, folder := range folders {
		if folder.IsDir() && utils.IncludeOrEmpty(folder.Name(), p.ids) {
			paths = append(paths, filepath.Join(folderPath, folder.Name()))
		}
	}

	p.dialogsPb.Finish()
	p.dialogsPb.Reset(len(paths))

	return paths, nil
}

func (p *Parser) loadAlbums() ([]string, error) {
	folderPath := path.Join(p.path, photos_dir, photo_albums_dir)
	albumsFiles, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, albumFile := range albumsFiles {
		albumId := utils.FileNameWithoutExtSliceNotation(albumFile.Name())
		if !albumFile.IsDir() && utils.IncludeOrEmpty(albumId, p.albumsIds) {
			paths = append(paths, filepath.Join(folderPath, albumFile.Name()))
		}
	}

	return paths, nil
}

func (p *Parser) StartParser(parentConext context.Context) {
	p.errGroup, _ = errgroup.WithContext(parentConext)
	p.errGroup.Go(p.parse)
}

func (p *Parser) parse() error {
	defer close(p.output)

	p.logger.Println("parser started")

	albumFiles, err := p.loadAlbums()

	if err != nil {
		p.logger.Printf("failed to load archive: %v", err)

		return fmt.Errorf("failed to load archive: %v", err)
	}

	p.logger.Printf("founded %d albums\n", len(albumFiles))

	p.albumsPb.SetTotal(len(albumFiles))

	for _, albumFile := range albumFiles {
		err := p.processAlbum(albumFile)
		p.albumsPb.Increment()
		if err != nil {
			p.logger.Printf("failed to parse album %s: %x", dir, err)

			return fmt.Errorf("failed to process album %s: %v", dir, err)
		}
	}
	p.albumsPb.Finish()

	return nil

	dirs, err := p.loadDialogs()
	if err != nil {
		p.logger.Printf("failed to load archive: %v", err)

		return fmt.Errorf("failed to load archive: %v", err)
	}

	p.logger.Printf("founded %d dialogs\n", len(dirs))

	p.dialogsPb.Finish()
	p.dialogsPb.Reset(len(dirs))

	for _, dir := range dirs {
		err := p.processDialog(dir)
		if err != nil {
			p.logger.Printf("failed to parse dialog %s: %x", dir, err)

			return fmt.Errorf("failed to process dialog %s: %v", dir, err)
		}
	}

	return nil
}
