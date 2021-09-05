package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils"
)

func (p *Parser) processFile(dialogName, filePath string) (functionError error) {
	defer utils.PanicInterceptor(&functionError)

	doc, err := parseFile(filePath, p.encoding)
	if err != nil {
		p.logger.Printf("error: failed to parse file %s\n", filePath)
		return err
	}

	doc.Find(".item .attachment").Each(func(i int, s *goquery.Selection) {
		attachemtType := s.Find(".attachment__description").First().Text()

		if utils.IncludeOrEmpty(attachemtType, p.types) {
			link := s.Find(".attachment__link").First()
			attachemtUrl, exist := link.Attr("href")
			if exist && len(attachemtUrl) > 0 {
				p.logger.Printf("founded `%s` in dialog `%s`: %s\n", attachemtType, dialogName, attachemtUrl)

				p.output <- contract.Attachemt{
					DialogName: dialogName,
					Url:        attachemtUrl,
					Type:       attachemtType,
					IsLink:     isLink(link) || isVideo(link),
				}

				p.attachemtPb.SetTotal(int(p.attachemtPb.Total) + 1)
			}
		}
	})

	return nil
}

func (p *Parser) getDialogName(filePath string) (string, error) {
	doc, err := parseFile(filePath, p.encoding)
	if err != nil {
		p.logger.Printf("error: failed to parse file %s\n", filePath)
		return "", err
	}

	crumbNode := doc.Find(".page_block_header .page_block_header_inner .ui_crumb").Last()

	return crumbNode.Text(), nil
}
