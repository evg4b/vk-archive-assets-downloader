package application

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cheggaaa/pb"
)

func (d *Downloader) Run(ctx context.Context) error {
	pool, err := pb.StartPool(d.dialogsPb, d.dialogPagesPb, d.attachemtPb)
	if err != nil {
		fmt.Fprintln(os.Stdout, "An error occurred while initializing the display of download processes.")
		fmt.Fprintln(os.Stdout, "The download process can be continued.")
		fmt.Fprintln(os.Stdout, "You can find information about the download status in the log file")
		fmt.Fprintf(os.Stdout, "Error information: %v\n", err)
		log.Print(err)
	} else {
		defer func() { _ = pool.Stop() }()
	}

	d.parser.StartParser(ctx)
	d.loader.StartLoading(ctx)

	err = d.parser.Wait()
	if err != nil {
		return err
	}

	return d.loader.Wait()
}
