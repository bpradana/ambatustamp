package common_test

import (
	"ambatustamp/common"
	"testing"
)

func Test_ChangeExtension(t *testing.T) {
	path := "path/to/directory/document.png"
	newExtension := ".pdf"
	newPath := common.ChangeExtension(path, newExtension)
	if newPath != "path/to/directory/document.pdf" {
		t.Errorf("expected %v, got %v", "path/to/directory/document.pdf", newPath)
	}
}
