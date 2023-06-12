package common

import (
	"path/filepath"
)

func GetAbsoluteFilePath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}
