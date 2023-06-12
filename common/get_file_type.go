package common

import (
	"path/filepath"
	"strings"
)

type FileType string

const (
	PDF      FileType = "pdf"
	Image    FileType = "image"
	Document FileType = "document"
)

func GetFileType(path string) FileType {
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".pdf":
		return PDF
	case ".jpg", ".jpeg", ".png":
		return Image
	case ".doc", ".docx":
		return Document
	default:
		return ""
	}
}
