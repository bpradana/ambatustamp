package common

import (
	"fmt"
	"os/exec"
)

func ConvertImageToPDF(imagePath string, pdfPath string) (string, error) {
	command := fmt.Sprintf("convert %s %s", imagePath, pdfPath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return pdfPath, nil
}

func ConvertDocumentToPDF(documentPath string, pdfPath string) (string, error) {
	command := fmt.Sprintf("soffice --headless --convert-to pdf %s --outdir %s", documentPath, pdfPath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return pdfPath, nil
}
