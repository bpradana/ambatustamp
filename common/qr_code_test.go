package common_test

import (
	"ambatustamp/common"
	"testing"
)

func Test_QrCode(t *testing.T) {
	outPath := "../assets/qr_code.png"
	logoPath := "../assets/logo.png"
	content := "https://www.google.com"
	size := 128

	_, err := common.GenerateQRCode(outPath, logoPath, content, size)
	if err != nil {
		t.Error(err)
	}
}
