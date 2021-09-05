package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils/files"
)

func (p *Parser) processFile(dialogName, filePath string) error {
	doc, err := files.ParseFile(filePath, p.encoding)
	if err != nil {
		p.log.Printf("ERROR: failed to parse file %s\n", filePath)
		return err
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
	doc, err := files.ParseFile(filePath, p.encoding)
	if err != nil {
		p.log.Printf("ERROR: failed to parse file %s\n", filePath)
		return "", err
	}

	crumbNode := doc.Find(".page_block_header .page_block_header_inner .ui_crumb").Last()

	return crumbNode.Text(), nil
}
