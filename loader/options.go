package loader

import "github.com/cheggaaa/pb"

func Destination(path string) LoaderOption {
	return func(loader *Loader) {
		loader.dest = path
	}
}

func AttachemtProgressBar(progressBar *pb.ProgressBar) LoaderOption {
	return func(loader *Loader) {
		loader.attachemtPb = progressBar
	}
}

func ThreadsCount(threadsCount int) LoaderOption {
	return func(loader *Loader) {
		loader.threadsCount = threadsCount
	}
}
