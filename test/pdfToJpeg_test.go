package test

import (
	"bytes"
	"fmt"
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

	t.Run("convert multiple pdf to jpeg", func(t *testing.T) {
		defer func() {
			require.NoError(t, os.RemoveAll("../images"))
		}()
		for i := range 10 {
			body := []byte(fmt.Sprintf(`{"filename": "test-%d"}`, i))
			res, err := http.Post(url, "application/json", bytes.NewReader(body))
			if err != nil {
				t.Error(err)
			}
			require.Equal(t, http.StatusOK, res.StatusCode)
		}
	})

	t.Run("convert inexistant pdf to jpeg", func(t *testing.T) {
		body := []byte(`{"filename": "___.pdf"}`)
		res, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}
		require.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("convert multiple inexistant pdf to jpeg", func(t *testing.T) {
		defer func() {
			require.NoError(t, os.RemoveAll("../images"))
		}()
		for i := range 10 {
			body := []byte(fmt.Sprintf(`{"filename": "%d.pdf"}`, i))
			res, err := http.Post(url, "application/json", bytes.NewReader(body))
			if err != nil {
				t.Error(err)
			}
			require.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	})
}
