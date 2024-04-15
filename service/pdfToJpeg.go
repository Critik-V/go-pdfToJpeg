package service

import (
	"errors"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

var ErrPickingPage = errors.New("error picking page")
var ErrCreatingJpeg = errors.New("error creating jpeg image")
var ErrEncodingJpeg = errors.New("error encoding jpeg image")

var ErrPdfDirNotExist = errors.New("pdf directory does not exist")
var ErrImgDirCreation = errors.New("image directory creation failed")

const jpegQuality int = 5    // Quality of the JPEG image
const imgExt string = ".jpg" // Extension of
const docExt string = ".pdf" // Extension of the document

func convert(doc *fitz.Document, imgDir string, fileName string) error {
	img, err := doc.Image(0)
	if err != nil {
		return errors.Join(ErrPickingPage, err)
	}

	f, err := os.Create(filepath.Join(imgDir, fmt.Sprintf("%v%v", fileName, imgExt)))
	if err != nil {
		return errors.Join(ErrCreatingJpeg, err)
	}

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: int(jpegQuality)})
	if err != nil {
		return errors.Join(ErrEncodingJpeg, err)
	}
	defer f.Close()
	return nil
}

func PdfToJpeg(fileName string) error {

	var pdfDir string = os.Getenv("PDF_STORAGE_PATH")   // Path to the directory where the PDFs are stored
	var imgDir string = os.Getenv("IMAGE_STORAGE_PATH") // Path to the directory where the images will be stored

	doc, err := fitz.New(fmt.Sprintf("%v/%v%v", pdfDir, fileName, docExt))
	if err != nil {
		return errors.Join(ErrPdfDirNotExist, err)
	}

	defer doc.Close()

	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		// img directory does not exist
		err = os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			return errors.Join(ErrImgDirCreation, err)
		}
		return convert(doc, imgDir, fileName)
	}
	// img directory already exists
	return convert(doc, imgDir, fileName)
}
