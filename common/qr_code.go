package common

import (
	"bytes"
	"image"
	"io/ioutil"

	"github.com/yeqown/go-qrcode"
)

func GenerateQRCode(outPath string, logoPath string, content string, size int) (string, error) {
	// load logo
	logoBytes, err := ioutil.ReadFile(logoPath)
	if err != nil {
		return "", err
	}

	decodedImage, _, err := image.Decode(bytes.NewReader(logoBytes))
	if err != nil {
		return "", err
	}

	qrLogo, err := qrcode.NewWithConfig(
		content,
		qrcode.DefaultConfig(),
		qrcode.WithBuiltinImageEncoder(qrcode.PNG_FORMAT),
		qrcode.WithLogoImage(decodedImage),
		qrcode.WithQRWidth(uint8(size)),
	)

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err := qrLogo.SaveTo(&buffer); err != nil {
		return "", err
	}

	err = ioutil.WriteFile(outPath, buffer.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	return outPath, nil
}
