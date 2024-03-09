package handlers

import (
	"go-pdf2jpeg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	FileName string `json:"fileName"`
}

func ConvertPdf(ctx *gin.Context) {
	var body body
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.PdfToJpeg(body.FileName)
	ctx.JSON(http.StatusOK, gin.H{"message": "Conversion successful"})
}
