package parser

import (
	"fmt"
	"os"
	"path/filepath"

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

	return nil, fmt.Errorf("unsupported encoding %s", encodingName)
}

func findDialogPages(dirPath string) (files []string, err error) {
	paths, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	files = []string{}
	for _, v := range paths {
		if v.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, v.Name())
		files = append(files, filePath)
	}

	return files, nil
}
