package loader

import (
	"log"
	"sync"

	"github.com/evg4b/vk-archive-assets-downloader/contract"
)

type Loader struct {
	Input <-chan contract.Attachemt
	Dest  string
	Wg    *sync.WaitGroup
	log   *log.Logger
}

func NewLoader(wg *sync.WaitGroup, dest string, input <-chan contract.Attachemt) *Loader {
	return &Loader{
		Input: input,
		Dest:  dest,
		Wg:    wg,
		log:   log.New(log.Writer(), "Loader |", log.Flags()),
	}
}
