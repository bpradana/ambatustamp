package common

import (
	"path/filepath"
)

func ChangeExtension(path string, newExtension string) string {
	extension := filepath.Ext(path)
	return path[0:len(path)-len(extension)] + newExtension
}
