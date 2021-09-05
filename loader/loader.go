package loader

import (
	"github.com/cheggaaa/pb"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"golang.org/x/sync/errgroup"
)

type LoaderOption = func(parser *Loader)

type Loader struct {
	input        <-chan contract.Attachemt
	dest         string
	threadsCount int
	attachemtPb  *pb.ProgressBar
	errGroup     *errgroup.Group
}

func NewLoader(input <-chan contract.Attachemt, options ...LoaderOption) *Loader {
	loader := &Loader{
		input:        input,
		dest:         "dest",
		attachemtPb:  pb.New(0),
		threadsCount: 10,
		errGroup:     &errgroup.Group{},
	}

	for _, option := range options {
		option(loader)
	}

	return loader
}

func (p *Loader) Wait() error {
	return p.errGroup.Wait()
}
