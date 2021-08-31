package common

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

func ParseFile(filePath string) (*goquery.Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("ERROR: failed to open %s\n", filePath)
		return nil, err
	}

	defer file.Close()

	decoder := charmap.Windows1251.NewDecoder()
	doc, err := goquery.NewDocumentFromReader(decoder.Reader(file))
	if err != nil {
		log.Printf("ERROR: failed to parse file %s\n", filePath)
		return nil, err
	}

	return doc, nil
}
