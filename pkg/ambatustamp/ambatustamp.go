package ambatustamp

import (
	"ambatustamp/common"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type StampConfig struct {
	Size     int
	LogoPath string
	Content  string
	Position string
	Xoffset  int
	Yoffset  int
}

type MetadataConfig struct {
	Title   string
	Author  string
	Subject string
}

type AmbatustampContract interface {
	LoadFile(filePath string) error
	Decrypt(password string) error
	Stamp(stampConfig *StampConfig) error
	Metadata(metadataConfig *MetadataConfig) error
	SaveFile(filePath string) error
}

type Ambatustamp struct {
	filePath            string
	isPasswordProtected bool
	password            string
}

func NewAmbatustamp() AmbatustampContract {
	return &Ambatustamp{
		filePath: "",
	}
}

func (a *Ambatustamp) LoadFile(filePath string) error {
	absPath := common.GetAbsoluteFilePath(filePath)
	log.Printf("[ambatustamp] [LoadFile] - loading file: %s", absPath)

	// Check file type
	switch common.GetFileType(absPath) {
	case common.PDF:
		// Check if PDF has password
		a.isPasswordProtected = common.PDFHasPassword(absPath)
		if a.isPasswordProtected {
			log.Printf("[ambatustamp] [LoadFile] - got PDF with password, please use .Decrypt(): %s", absPath)
		}

		// Copy file to temp file
		log.Printf("[ambatustamp] [LoadFile] - got PDF, creating temp file: %s", absPath)
		a.filePath = common.ChangeExtension(absPath, fmt.Sprintf("_%d.pdf", time.Now().Unix()))
		command := fmt.Sprintf("cp %s %s", absPath, a.filePath)
		cmd := exec.Command("sh", "-c", command)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("[ambatustamp] [LoadFile] - error creating temp file: %s", err)
			return err
		}
		log.Printf("[ambatustamp] [LoadFile] - temp file: %s", a.filePath)
	case common.Image:
		// Convert image to PDF
		log.Printf("[ambatustamp] [LoadFile] - got image, converting to PDF: %s", absPath)
		pdfPath := common.ChangeExtension(absPath, fmt.Sprintf("_%d.pdf", time.Now().Unix()))
		_, err := common.ConvertImageToPDF(absPath, pdfPath)
		if err != nil {
			log.Fatalf("[ambatustamp] [LoadFile] - error converting image to PDF: %s", err)
			return err
		}
		a.filePath = pdfPath
		log.Printf("[ambatustamp] [LoadFile] - converted image to PDF: %s", a.filePath)
	case common.Document:
		// Convert document to PDF
		pdfPath := common.ChangeExtension(absPath, fmt.Sprintf("_%d.pdf", time.Now().Unix()))
		_, err := common.ConvertDocumentToPDF(absPath, pdfPath)
		if err != nil {
			return err
		}
		a.filePath = pdfPath
		log.Printf("[ambatustamp] [LoadFile] - converted document to PDF: %s", a.filePath)
	default:
		log.Fatalf("[ambatustamp] [LoadFile] - invalid file type: %s", absPath)
		return errors.New("invalid file type")
	}
	return nil
}

func (a *Ambatustamp) Decrypt(password string) error {
	if a.filePath == "" {
		log.Fatalf("[ambatustamp] [Stamp] - no file loaded")
		return errors.New("no file loaded")
	}

	a.password = password

	// Decrypt PDF
	command := fmt.Sprintf("pdfcpu decrypt -upw %s %s", password, a.filePath)
	log.Printf("[ambatustamp] [Decrypt] - decrypting PDF: %s", a.filePath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [Decrypt] - error decrypting PDF: %s", err)
		return err
	}

	return nil
}

func (a *Ambatustamp) Stamp(stampConfig *StampConfig) error {
	if a.filePath == "" {
		log.Fatalf("[ambatustamp] [Stamp] - no file loaded")
		return errors.New("no file loaded")
	}

	// Generate QR code
	absLogoPath := common.GetAbsoluteFilePath(stampConfig.LogoPath)
	outPath := common.ChangeExtension(absLogoPath, "_qr.png")
	content := stampConfig.Content
	qrCodePath, err := common.GenerateQRCode(outPath, absLogoPath, content, stampConfig.Size)
	log.Printf("[ambatustamp] [Stamp] - generating QR code: %s", qrCodePath)
	if err != nil {
		log.Fatalf("[ambatustamp] [Stamp] - error generating QR code: %s", err)
		return err
	}

	// Stamp PDF
	command := fmt.Sprintf("pdfcpu stamp add -pages even,odd -mode image -- '%s' 'pos:%s, rot:0, sc:.1, offset:%d %d' %s", qrCodePath, stampConfig.Position, stampConfig.Xoffset, stampConfig.Yoffset, a.filePath)
	log.Printf("[ambatustamp] [Stamp] - stamping PDF: %s", a.filePath)
	cmd := exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [Stamp] - error stamping PDF: %s", err)
		return err
	}

	// Remove QR code
	command = fmt.Sprintf("rm %s", qrCodePath)
	log.Printf("[ambatustamp] [Stamp] - removing QR code: %s", qrCodePath)
	cmd = exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [Stamp] - error removing QR code: %s", err)
		return err
	}

	return nil
}

func (a *Ambatustamp) Metadata(metadataConfig *MetadataConfig) error {
	if a.filePath == "" {
		log.Fatalf("[ambatustamp] [Stamp] - no file loaded")
		return errors.New("no file loaded")
	}

	// Add metadata
	command := fmt.Sprintf("pdfcpu properties add %s 'Title = %s' 'Author = %s' 'Subject = %s'", a.filePath, metadataConfig.Title, metadataConfig.Author, metadataConfig.Subject)
	log.Printf("[ambatustamp] [Metadata] - adding metadata to PDF: %s", a.filePath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [Metadata] - error adding metadata to PDF: %s", err)
		return err
	}

	return nil
}

func (a *Ambatustamp) SaveFile(filePath string) error {
	if a.filePath == "" {
		log.Fatalf("[ambatustamp] [Stamp] - no file loaded")
		return errors.New("no file loaded")
	}

	// If file is password protected, re-lock it
	if a.isPasswordProtected {
		command := fmt.Sprintf("pdfcpu encrypt -upw %s -opw %s %s", a.password, a.password, a.filePath)
		log.Printf("[ambatustamp] [SaveFile] - re-locking file: %s", filePath)
		cmd := exec.Command("sh", "-c", command)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("[ambatustamp] [SaveFile] - error re-locking file: %s", err)
			return err
		}
	}
	command := fmt.Sprintf("cp %s %s", a.filePath, filePath)
	log.Printf("[ambatustamp] [SaveFile] - saving file: %s", filePath)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [SaveFile] - error saving file: %s", err)
		return err
	}

	// Remove temp file
	command = fmt.Sprintf("rm %s", a.filePath)
	log.Printf("[ambatustamp] [SaveFile] - removing temp file: %s", a.filePath)
	cmd = exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("[ambatustamp] [SaveFile] - error removing temp file: %s", err)
		return err
	}

	return nil
}
