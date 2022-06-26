package loader

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/evg4b/vk-archive-assets-downloader/contract"
)

func getLogger(index int) *log.Logger {
	return log.New(log.Writer(), fmt.Sprintf("[loader thread %v] ", index+1), log.Flags())
}

func getFilePath(dest string, attachemt contract.Attachemt) (string, string, error) {
	parserUrl, err := url.Parse(attachemt.Url)
	if err != nil {
		return "", "", err
	}
	fileName := filepath.Base(parserUrl.Path)
	var directoryPath string

	switch attachemt.Section {
	case contract.Dialog:
		directoryPath = filepath.Join(dest, attachemt.Name, attachemt.Type)
	case contract.Album:
		directoryPath = filepath.Join(dest, "Album["+attachemt.Name+"]")
	default:
		return "", "", fmt.Errorf("unknown attachmentType %d", attachemt.Section)
	}

	err = os.MkdirAll(directoryPath, os.ModePerm)
	if err != nil {
		return "", "", err
	}

	return directoryPath, fileName, nil
}
