package test

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

const url string = "http://localhost:5001/convert"

func TestConvertPdfToJpeg(t *testing.T) {

	// t.Run("convert one pdf to jpeg", func(t *testing.T) {
	// 	defer func() {
	// 		require.NoError(t, os.RemoveAll("../images"))
	// 	}()
	// 	body := []byte(`{"filename": "test"}`)
	// 	res, err := http.Post(url, "application/json", bytes.NewReader(body))
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	require.Equal(t, http.StatusOK, res.StatusCode)
	// })

	t.Run("convert multiple pdf to jpeg", func(t *testing.T) {
		defer func() {
			require.NoError(t, os.RemoveAll("../images"))
		}()
		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				body := []byte(fmt.Sprintf(`{"filename": "test-%d"}`, i))
				res, err := http.Post(url, "application/json", bytes.NewReader(body))
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, http.StatusOK, res.StatusCode)
			}()
		}
		wg.Wait()
	})
	// t.Run("convert inexistent pdf to jpeg", func(t *testing.T) {
	// 	body := []byte(`{"filename": "___.pdf"}`)
	// 	res, err := http.Post(url, "application/json", bytes.NewReader(body))
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	// })

	// t.Run("convert multiple inexistent pdf to jpeg", func(t *testing.T) {
	// 	defer func() {
	// 		require.NoError(t, os.RemoveAll("../images"))
	// 	}()
	// 	for i := range 5 {
	// 		body := []byte(fmt.Sprintf(`{"filename": "%d.pdf"}`, i))
	// 		res, err := http.Post(url, "application/json", bytes.NewReader(body))
	// 		if err != nil {
	// 			t.Error(err)
	// 		}
	// 		require.Equal(t, http.StatusBadRequest, res.StatusCode)
	// 	}
	// })
}
