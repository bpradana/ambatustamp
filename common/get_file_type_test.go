package common_test

import (
	"ambatustamp/common"
	"testing"
)

func Test_GetFileType(t *testing.T) {
	path := "path/to/directory/document.png"
	fileType := common.GetFileType(path)
	if fileType != common.Image {
		t.Errorf("expected %v, got %v", common.Image, fileType)
	}

	path = "path/to/directory/document.pdf"
	fileType = common.GetFileType(path)
	if fileType != common.PDF {
		t.Errorf("expected %v, got %v", common.PDF, fileType)
	}

	path = "path/to/directory/document.doc"
	fileType = common.GetFileType(path)
	if fileType != common.Document {
		t.Errorf("expected %v, got %v", common.Document, fileType)
	}
}
