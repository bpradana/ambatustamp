package ambatustamp_test

import (
	"ambatustamp/pkg/ambatustamp"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func Test_Ambatustamp_PDF(t *testing.T) {
	amb := ambatustamp.NewAmbatustamp()
	if amb == nil {
		t.Error("expected ambatustamp to not be nil")
	}

	filePath := "../../assets/test_pdf.pdf"
	err := amb.LoadFile(filePath)
	if err != nil {
		t.Error(err)
	}

	stampConfig := ambatustamp.StampConfig{
		Size:     30,
		LogoPath: "../../assets/logo.png",
		Content:  uuid.New().String(),
		Position: "br",
		Xoffset:  -25,
		Yoffset:  25,
	}
	err = amb.Stamp(&stampConfig)
	fmt.Print(err)
	if err != nil {
		t.Error(err)
	}

	metadataConfig := ambatustamp.MetadataConfig{
		Title:   "Test PDF",
		Author:  "Ambatustamp",
		Subject: "Test PDF",
	}
	err = amb.Metadata(&metadataConfig)
	if err != nil {
		t.Error(err)
	}

	outPath := "../../assets/test_pdf_stamped.pdf"
	err = amb.SaveFile(outPath)
	if err != nil {
		t.Error(err)
	}
}

func Test_Ambatustamp_Image(t *testing.T) {
	amb := ambatustamp.NewAmbatustamp()
	if amb == nil {
		t.Error("expected ambatustamp to not be nil")
	}

	filePath := "../../assets/test_image.png"
	err := amb.LoadFile(filePath)
	if err != nil {
		t.Error(err)
	}

	stampConfig := ambatustamp.StampConfig{
		Size:     64,
		LogoPath: "../../assets/logo.png",
		Content:  uuid.New().String(),
		Position: "br",
		Xoffset:  -25,
		Yoffset:  25,
	}
	err = amb.Stamp(&stampConfig)
	fmt.Print(err)
	if err != nil {
		t.Error(err)
	}

	metadataConfig := ambatustamp.MetadataConfig{
		Title:   "Test PDF",
		Author:  "Ambatustamp",
		Subject: "Test PDF",
	}
	err = amb.Metadata(&metadataConfig)
	if err != nil {
		t.Error(err)
	}

	outPath := "../../assets/test_pdf_stamped.pdf"
	err = amb.SaveFile(outPath)
	if err != nil {
		t.Error(err)
	}
}
