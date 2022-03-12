package parser

import (
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils"
)

func (p *Parser) processAlbum(file string) (functionError error) {
	defer utils.PanicInterceptor(&functionError)

	err := p.processAlbumFile(file)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) processAlbumFile(filePath string) (functionError error) {
	defer utils.PanicInterceptor(&functionError)

	albumId := utils.FileNameWithoutExtSliceNotation(filepath.Base(filePath))

	doc, err := parseFile(filePath, p.encoding)
	if err != nil {
		p.logger.Printf("error: failed to parse file %s\n", filePath)
		return err
	}

	doc.Find(".item__main > a > img").Each(func(i int, s *goquery.Selection) {
		photoId, exists := s.Attr("alt")

		if exists && len(photoId) > 0 {
			photoSrc, exist := s.Attr("src")
			if exist && len(photoSrc) > 0 {
				p.logger.Printf("founded src in album `%s`: %s\n", albumId, photoSrc)

				p.output <- contract.Attachemt{
					AttachmentType: contract.Album,
					Name:           albumId,
					Url:            photoSrc,
					Type:           "Image",
					IsLink:         false,
				}

				p.attachemtPb.SetTotal(int(p.attachemtPb.Total) + 1)
			}
		}

	})

	return nil
}
