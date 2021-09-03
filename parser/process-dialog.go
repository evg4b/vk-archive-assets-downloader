package parser

import (
	"os"
	"path/filepath"
)

func (p *Parser) parseDialog(dirPath string) ([]string, error) {
	p.log.Printf("started working in directory: %s\n", dirPath)
	paths, err := os.ReadDir(dirPath)
	if err != nil {
		p.log.Printf("ERROR: failed to read dir %s\n", dirPath)
		return nil, err
	}

	files := []string{}
	for _, v := range paths {
		if v.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, v.Name())
		p.log.Printf("Founded file %s\n", filePath)
		files = append(files, filePath)
	}

	return files, nil
}
