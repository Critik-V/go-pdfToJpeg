package handlers

import (
	"errors"
	"go-pdf2jpeg/service"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type body struct {
	FileName string `json:"filename"`
}

func ConvertPdf(ctx *gin.Context) {
	var body body
	finished := make(chan bool)
	mutex := &sync.Mutex{}

	mutex.Lock()
	go func() {
		defer mutex.Unlock()
		err := ctx.BindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		err = service.PdfToJpeg(body.FileName)
		if err == nil {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Conversion successful", "status": "success"})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("conversion failed").Error()})
		}
		finished <- true
	}()
	<-finished
}
