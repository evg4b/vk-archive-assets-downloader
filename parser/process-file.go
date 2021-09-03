package parser

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils/files"
	"golang.org/x/text/encoding/charmap"
)

func (p *Parser) processFile(dialogName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		p.log.Printf("ERROR: failed to open %s: %x\n", filePath, err)
		return err
	}

	defer file.Close()

	decoder := charmap.Windows1251.NewDecoder()
	doc, err := goquery.NewDocumentFromReader(decoder.Reader(file))
	if err != nil {
		if err != nil {
			p.log.Printf("ERROR: failed to parse file %s\n", filePath)
			return err
		}
	}

	doc.Find(".item .attachment").Each(func(i int, s *goquery.Selection) {
		description := s.Find(".attachment__description").First()
		link := s.Find(".attachment__link").First()
		linkAdders, exist := link.Attr("href")
		if exist && len(linkAdders) > 0 {
			p.log.Printf("Founded attachment %s\n", linkAdders)
			p.output <- contract.Attachemt{
				DialogName: dialogName,
				Url:        linkAdders,
				Type:       strings.Trim(description.Text(), " "),
			}
		}
	})

	return nil
}

func (p *Parser) getDialogName(filePath string) (string, error) {
	doc, err := files.ParseFile(filePath)
	if err != nil {
		return "", err
	}

	crumbNode := doc.Find(".page_block_header .page_block_header_inner .ui_crumb").Last()

	return crumbNode.Text(), nil
}
