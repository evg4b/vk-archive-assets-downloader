package common

import "github.com/cheggaaa/pb"

func InitProgressBar(progressBar *pb.ProgressBar, count int) {
	if progressBar != nil {
		progressBar.Reset(count)
	}
}

func IncrementTotal(progressBar *pb.ProgressBar, count int) {
	if progressBar != nil {
		progressBar.SetTotal64(progressBar.Total + 1)
	}
}
