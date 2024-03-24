package service

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

const jpegQuality int = 7    // Quality of the JPEG image
const imgExt string = ".jpg" // Extension of
const docExt string = ".pdf" // Extension of the document

func convert(doc *fitz.Document, imgDir string, fileName string) {
	img, err := doc.Image(0)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join(imgDir, fmt.Sprintf("%v%v", fileName, imgExt)))
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: int(jpegQuality)})
	if err != nil {
		panic(err)
	}

	f.Close()
}

func PdfToJpeg(fileName string) {

	var pdfDir string = os.Getenv("PDF_STORAGE_PATH")   // Path to the directory where the PDFs are stored
	var imgDir string = os.Getenv("IMAGE_STORAGE_PATH") // Path to the directory where the images will be stored

	doc, err := fitz.New(fmt.Sprintf("%v/%v%v", pdfDir, fileName, docExt))
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		// img directory does not exist
		err = os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		convert(doc, imgDir, fileName)
	} else {
		// img directory already exists
		convert(doc, imgDir, fileName)
	}
}
