package common

import (
	"fmt"
	"os/exec"
)

func PDFHasPassword(pdfPath string) bool {
	command := fmt.Sprintf("pdfcpu validate %s", pdfPath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	return err != nil
}
