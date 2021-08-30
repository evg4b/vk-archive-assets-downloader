package loader

import (
	"context"
	"log"
	"sync"

	"github.com/evg4b/vk-archive-assets-downloader/internal/common"
)

type Loader struct {
	Input <-chan common.Attachemt
	Dest  string
	Wg    *sync.WaitGroup
}

func (p *Loader) Load(ctx context.Context) {
	defer p.Wg.Done()
	log.Println("Loader started")

	for v := range p.Input {
		log.Print(v)
	}
}
