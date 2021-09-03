package loader

import (
	"context"
)

func (p *Loader) Load(ctx context.Context) {
	defer p.Wg.Done()
	p.log.Println("Loader started")

	for v := range p.Input {
		p.log.Print(v)
	}
}
