package service

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

const jpegQuality int = 15  // Quality of the JPEG image
const dir string = "images" // Directory to save images
const ext string = ".jpg"   // Extension of the image

func PdfToJpeg(fileName string) {
	doc, err := fitz.New("pdf/test.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		img, err := doc.Image(0)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(dir, fmt.Sprintf("%v%v", fileName, ext)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: int(jpegQuality)})
		if err != nil {
			panic(err)
		}

		f.Close()
	} else {
		fmt.Println("Directory already exists")
	}
}
