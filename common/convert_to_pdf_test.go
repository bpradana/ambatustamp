package common_test

import (
	"ambatustamp/common"
	"testing"
)

func Test_ConvertImageToPDF(t *testing.T) {
	imagePath := "../assets/logo.png"
	pdfPath := "../assets/logo.pdf"
	_, err := common.ConvertImageToPDF(imagePath, pdfPath)
	if err != nil {
		t.Error(err)
	}
}

func Test_ConvertDocumentToPDF(t *testing.T) {
	documentPath := "../assets/test_document.docx"
	pdfPath := "../assets/test_document.pdf"
	_, err := common.ConvertDocumentToPDF(documentPath, pdfPath)
	if err != nil {
		t.Error(err)
	}
}
