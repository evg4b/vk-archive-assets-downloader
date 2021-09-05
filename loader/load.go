package loader

import (
	"context"

	"github.com/evg4b/vk-archive-assets-downloader/utils"
	"github.com/hashicorp/go-multierror"
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

func (p *Loader) loadingThread(index int, ctx context.Context) (threadError error) {
	defer utils.PanicInterceptor(&threadError)

	logger := getLogger(index)
	logger.Println("loading thread started")

	var result *multierror.Error

	for attachemt := range p.input {
		err := p.downloadFile(logger, attachemt)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result.ErrorOrNil()
}
