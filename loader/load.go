package loader

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func (p *Loader) StartLoading(parentConext context.Context) {
	errGroup, ctx := errgroup.WithContext(parentConext)
	p.errGroup = errGroup

	for i := 0; i < p.threadsCount; i++ {
		index := i
		p.errGroup.Go(func() error {
			return p.loadingThread(index, ctx)
		})
	}
}

func (p *Loader) loadingThread(index int, ctx context.Context) error {
	logger := getLogger(index)
	logger.Println("loading thread started")

	for v := range p.input {
		logger.Printf("%s:%s %s", v.DialogName, v.Type, v.Url)
		p.attachemtPb.Increment()
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func getLogger(index int) *log.Logger {
	return log.New(log.Writer(), fmt.Sprintf("[loader thread %v] ", index+1), log.Flags())
}
