package application

import (
	"context"

	"github.com/cheggaaa/pb"
)

func (d *Downloader) Run(ctx context.Context) error {
	pool, err := pb.StartPool(d.dialogsPb, d.dialogPagesPb, d.attachemtPb)
	if err != nil {
		return err
	}

	defer pool.Stop()

	d.wg.Add(2)
	go d.parser.Parse(ctx)
	go d.loader.Load(ctx)

	d.wg.Wait()

	return nil
}
