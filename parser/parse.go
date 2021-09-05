package parser

import (
	"context"
	"fmt"
)

func (p *Parser) StartParser(ctx context.Context) {
	p.wg.Add(1)
	go p.parse(ctx)
}

func (p *Parser) parse(ctx context.Context) {
	defer p.wg.Done()
	defer close(p.output)

	p.log.Println("Parser started")

	dirs, err := p.load()
	if err != nil {
		panic(err)
	}

	p.log.Printf("Founded %d dialogs\n", len(dirs))

	p.dialogsPb.Finish()
	p.dialogsPb.Reset(len(dirs))

	for _, dir := range dirs {
		files, err := p.parseDialog(dir)
		if err != nil {
			p.log.Printf("ERROR: failed to read dialog %s\n", dir)
			continue
		}

		dialogName, err := p.getDialogName(files[0])
		if err != nil {
			p.log.Printf("ERROR: failed to read dis %s\n", dir)
			continue
		}

		p.log.Printf("Founded %s dialog name for path %s", dialogName, dir)

		p.dialogPagesPb.Prefix(fmt.Sprintf("Dialog with %s", dialogName))
		p.dialogPagesPb.Finish()
		p.dialogPagesPb.Reset(len(files))

		for _, filePath := range files {
			p.processFile(dialogName, filePath)
			p.dialogPagesPb.Increment()
		}

		p.dialogsPb.Increment()
	}
}
