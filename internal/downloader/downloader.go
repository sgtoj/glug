package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url, dir string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status returned (%d) for download request: %s", res.StatusCode, url)
	}

	fileName := filepath.Base(res.Request.URL.Path)
	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create local file for downloaded file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write local file for downloaded file")
	}

	return filePath, nil
}
