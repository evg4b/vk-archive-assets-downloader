package parser

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/internal/common"
	"golang.org/x/text/encoding/charmap"
)

const dir = "messages"

type Parser struct {
	Path   string
	Ids    []string
	Output chan<- common.Attachemt
	Wg     *sync.WaitGroup
}

func (p *Parser) Parse(ctx context.Context, dirs []string) {
	defer p.Wg.Done()
	defer close(p.Output)

	log.Println("Parser started")

	for _, dirPath := range dirs {
		log.Printf("Started working in directory: %s\n", dirPath)
		files, err := os.ReadDir(dirPath)
		if err != nil {
			log.Printf("ERROR: failed to read dis %s\n", dirPath)
			continue
		}

		for _, v := range files {
			if v.IsDir() {
				continue
			}

			log.Printf("Founded file %s\n", v.Name())
			filePath := filepath.Join(dirPath, v.Name())
			file, err := os.Open(filePath)
			if err != nil {
				log.Printf("ERROR: failed to open %s\n", filePath)
				continue
			}

			decoder := charmap.Windows1251.NewDecoder()
			doc, err := goquery.NewDocumentFromReader(decoder.Reader(file))
			if err != nil {
				if err != nil {
					log.Printf("ERROR: failed to parse file %s\n", filePath)
					continue
				}
			}

			doc.Find(".item .attachment").Each(func(i int, s *goquery.Selection) {
				description := s.Find(".attachment__description").First()
				link := s.Find(".attachment__link").First()
				linkAdders, exist := link.Attr("href")
				if len(linkAdders) > 0 && exist {
					log.Printf("Founded attachment %s\n", linkAdders)
					p.Output <- common.Attachemt{
						DialogName: "test",
						Url:        linkAdders,
						Type:       strings.Trim(description.Text(), " "),
					}
				}
			})

			file.Close()
		}
	}
}

func (p *Parser) Load(ctx context.Context) ([]string, error) {
	folderPath := path.Join(p.Path, dir)
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, folder := range folders {
		if folder.IsDir() && common.IncludeOrEmpty(folder.Name(), p.Ids) {
			paths = append(paths, filepath.Join(folderPath, folder.Name()))
		}
	}

	return paths, nil
}
