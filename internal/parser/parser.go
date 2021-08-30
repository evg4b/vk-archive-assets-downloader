package parser

import (
	"context"

	"github.com/evg4b/vk-archive-assets-downloader/internal"
)

type Parser struct {
	Path string
	Ids  []string
}

func (p *Parser) Run(ctx context.Context, output chan<- internal.Attachemt) {
}
