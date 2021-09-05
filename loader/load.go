package loader

import (
	"context"
)

func (p *Loader) StartLoading(ctx context.Context) {
	p.wg.Add(1)
	go p.loadingThread(ctx)
}

func (p *Loader) loadingThread(ctx context.Context) {
	defer p.wg.Done()
	p.log.Println("Loader started")

	for v := range p.input {
		p.log.Printf("%s:%s %s", v.DialogName, v.Type, v.Url)
	}
}
