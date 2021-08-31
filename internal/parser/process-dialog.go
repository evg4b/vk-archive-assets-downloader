package parser

import (
	"log"
	"os"
	"path/filepath"
)

func (p *Parser) parseDialog(dirPath string) ([]string, error) {
	log.Printf("Started working in directory: %s\n", dirPath)
	paths, err := os.ReadDir(dirPath)
	if err != nil {
		log.Printf("ERROR: failed to read dis %s\n", dirPath)
		return nil, err
	}

	files := []string{}
	for _, v := range paths {
		if v.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, v.Name())
		log.Printf("Founded file %s\n", filePath)
		files = append(files, filePath)
	}

	return files, nil
}
