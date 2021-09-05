package parser

import (
	"fmt"

	"github.com/evg4b/vk-archive-assets-downloader/utils"
)

func (p *Parser) processDialog(dir string) (functionError error) {
	defer utils.PanicInterceptor(&functionError)

	files, err := findDialogPages(dir)
	if err != nil {
		p.logger.Printf("error: failed to read dialog %s\n", dir)
		return err
	}

	p.logger.Printf("founded %d pages in dialog %s\n", len(files), dir)

	dialogName, err := p.getDialogName(files[0])
	if err != nil {
		p.logger.Printf("error: failed to read dis %s\n", dir)
		return err
	}

	p.logger.Printf("founded %s dialog name for path %s", dialogName, dir)

	p.dialogPagesPb.Prefix(fmt.Sprintf("Dialog with %s", dialogName))
	p.dialogPagesPb.Finish()
	p.dialogPagesPb.Reset(len(files))

	for _, filePath := range files {
		err := p.processFile(dialogName, filePath)
		if err != nil {
			return err
		}

		p.dialogPagesPb.Increment()
	}

	p.dialogsPb.Increment()

	return nil
}
