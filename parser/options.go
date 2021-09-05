package parser

import (
	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/utils/collections"
)

func Source(path string) ParserOption {
	return func(parser *Parser) {
		parser.path = path
	}
}

func AttachemtProgressBar(progressBar *pb.ProgressBar) ParserOption {
	return func(parser *Parser) {
		parser.attachemtPb = progressBar
	}
}

func DialogsProgressBar(progressBar *pb.ProgressBar) ParserOption {
	return func(parser *Parser) {
		parser.dialogsPb = progressBar
	}
}

func DialogPagesProgressBar(progressBar *pb.ProgressBar) ParserOption {
	return func(parser *Parser) {
		parser.dialogPagesPb = progressBar
	}
}

func Encoding(encoding string) ParserOption {
	return func(parser *Parser) {
		parser.encoding = encoding
	}
}

func Dialogs(dialogs string) ParserOption {
	return func(parser *Parser) {
		parser.ids = collections.SplitNotEmpty(dialogs)
	}
}

func Types(types string) ParserOption {
	return func(parser *Parser) {
		parser.types = collections.SplitNotEmpty(types)
	}
}
