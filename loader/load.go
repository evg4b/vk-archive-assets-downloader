package loader

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (p *Loader) StartLoading(ctx context.Context) {
	p.wg.Add(p.threadsCount)
	for i := 0; i < p.threadsCount; i++ {
		go p.loadingThread(i, ctx)
	}
}

func (p *Loader) loadingThread(index int, ctx context.Context) {
	defer p.wg.Done()
	logger := getLogger(index)
	logger.Println("Loading thread started")

	for v := range p.input {
		logger.Printf("%s:%s %s", v.DialogName, v.Type, v.Url)
		p.attachemtPb.Increment()
		time.Sleep(200 * time.Millisecond)
	}
}

func getLogger(index int) *log.Logger {
	return log.New(log.Writer(), fmt.Sprintf("Loader thread %v |", index+1), log.Flags())
}
