package main

import (
	"ambatustamp/pkg/ambatustamp"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	stampConfig := ambatustamp.StampConfig{
		Size:     128,
		LogoPath: "./assets/logo.png",
		Content:  uuid.New().String(),
		Position: "tr",
		Xoffset:  0,
		Yoffset:  0,
	}

	metadataConfig := ambatustamp.MetadataConfig{
		Title:   "Test PDF",
		Author:  "Ambatustamp",
		Subject: "Testing Ambatustamp",
	}

	amb := ambatustamp.NewAmbatustamp()
	err := amb.LoadFile("./assets/test_pdf.pdf")
	if err != nil {
		fmt.Print(err)
	}

	// uncomment this line if your pdf is password protected
	// err = amb.Decrypt("YOUR_PASSWORD")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	err = amb.Stamp(&stampConfig)
	if err != nil {
		fmt.Print(err)
	}

	err = amb.Metadata(&metadataConfig)
	if err != nil {
		fmt.Print(err)
	}

	outPath := "./assets/test_pdf_stamped.pdf"
	err = amb.SaveFile(outPath)
	if err != nil {
		fmt.Print(err)
	}
}
