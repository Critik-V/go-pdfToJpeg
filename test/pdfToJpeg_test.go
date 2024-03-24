package test

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const url string = "http://localhost:5001/convert"

func TestConvertPdfToJpeg(t *testing.T) {
	t.Run("convert one pdf to jpeg", func(t *testing.T) {
		defer func() {
			require.NoError(t, os.RemoveAll("../images"))
		}()
		body := []byte(`{"filename": "test"}`)
		res, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}
		require.Equal(t, http.StatusOK, res.StatusCode)
	})
}
