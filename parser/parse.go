package parser

import (
	"context"

	"github.com/evg4b/vk-archive-assets-downloader/utils/progressbar"
)

func (p *Parser) Parse(ctx context.Context) {
	defer p.wg.Done()
	defer close(p.output)

	p.log.Println("Parser started")

	dirs, err := p.load()
	if err != nil {
		panic(err)
	}

	p.log.Printf("Founded %d dialogs\n", len(dirs))

	for _, dirPath := range dirs {
		files, err := p.parseDialog(dirPath)
		if err != nil {
			p.log.Printf("ERROR: failed to read dis %s\n", dirPath)
			continue
		}

		progressbar.InitProgressBar(p.dialogPagesPb, len(files))
		dialogName, err := p.getDialogName(files[0])
		if err != nil {
			p.log.Printf("ERROR: failed to read dis %s\n", dirPath)
			continue
		}

		for _, filePath := range files {
			p.processFile(dialogName, filePath)
		}
	}
}
