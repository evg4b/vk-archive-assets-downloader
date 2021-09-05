package application

import (
	"context"
)

func (d *Downloader) Run(ctx context.Context) error {
	// pool, err := pb.StartPool(d.dialogsPb, d.dialogPagesPb, d.attachemtPb)
	// if err != nil {
	// 	return err
	// }

	// defer pool.Stop()

	d.parser.StartParser(ctx)
	d.loader.StartLoading(ctx)

	d.parser.Wait()
	d.loader.Wait()

	return nil
}
