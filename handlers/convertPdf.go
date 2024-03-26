package handlers

import (
	"go-pdf2jpeg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	FileName string `json:"filename"`
}

func ConvertPdf(ctx *gin.Context) {
	var body body
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	done := make(chan bool)

	go func() {
		err := service.PdfToJpeg(body.FileName)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			done <- true
		}
		done <- true
	}()

	<-done
	ctx.JSON(http.StatusOK, gin.H{"message": "Conversion successful"})
}
