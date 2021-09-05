package loader

import (
	"log"
	"sync"

	"github.com/evg4b/vk-archive-assets-downloader/contract"
)

type Loader struct {
	input <-chan contract.Attachemt
	dest  string
	wg    *sync.WaitGroup
	log   *log.Logger
}

func NewLoader(dest string, input <-chan contract.Attachemt) *Loader {
	return &Loader{
		wg:    &sync.WaitGroup{},
		input: input,
		dest:  dest,
		log:   log.New(log.Writer(), "Loader |", log.Flags()),
	}
}

func (p *Loader) Wait() {
	p.wg.Wait()
}
