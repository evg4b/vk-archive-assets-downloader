package loader

import (
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
)

type LoaderOption = func(parser *Loader)

type Loader struct {
	input        <-chan contract.Attachemt
	dest         string
	wg           *sync.WaitGroup
	threadsCount int
	attachemtPb  *pb.ProgressBar
}

func NewLoader(input <-chan contract.Attachemt, options ...LoaderOption) *Loader {
	loader := &Loader{
		wg:           &sync.WaitGroup{},
		input:        input,
		dest:         "dest",
		attachemtPb:  pb.New(0),
		threadsCount: 10,
	}

	if options != nil {
		for _, option := range options {
			option(loader)
		}
	}

	return loader
}

func (p *Loader) Wait() {
	p.wg.Wait()
}
