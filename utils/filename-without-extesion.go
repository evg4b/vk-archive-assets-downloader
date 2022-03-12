package utils

import "path/filepath"

func FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
