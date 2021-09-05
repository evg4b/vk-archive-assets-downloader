package parser

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func parseFile(filePath, encodingName string) (*goquery.Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if encoding, ok := encodings[encodingName]; ok {
		decoder := encoding.NewDecoder()
		doc, err := goquery.NewDocumentFromReader(decoder.Reader(file))
		if err != nil {
			return nil, err
		}

		return doc, nil
	}

	return nil, fmt.Errorf("Unsupported encoding %s", encodingName)
}
