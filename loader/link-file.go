package loader

import (
	"fmt"
	"os"
)

var fileContent string = `<html><head><meta http-equiv="refresh" content="0;url=%s"></head></html>`

func makeLinkFile(filename, url string) error {
	return os.WriteFile(
		fmt.Sprintf("%s-link.html", filename),
		[]byte(fmt.Sprintf(fileContent, url)),
		os.ModePerm,
	)
}
