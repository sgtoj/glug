package downloader

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sgtoj/glug/internal/registry"
)

var (
	ErrBinaryNotFound = errors.New("unable to find binary file")
)

func Move(srcDir, dstDir string, tool registry.ToolData) (string, error) {
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create destination directory for file: %w", err)
	}

	srcPath, err := findBinaryPath(srcDir, tool)
	if err != nil {
		return "", err
	}

	dstPath := filepath.Join(dstDir, tool.Name)
	if runtime.GOOS == "windows" {
		dstPath = dstPath + ".exe"
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", fmt.Errorf("unable to read binary file from temp directory: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return "", fmt.Errorf("unable to create binary file: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", fmt.Errorf("unable to copy binary file from temp directory: %w", err)
	}

	return dstPath, nil
}

// todo: need to find a better solution; this feels very sloppy
func findBinaryPath(srcDir string, tool registry.ToolData) (string, error) {
	// check binary is just named its name
	candidatePath := filepath.Join(srcDir, tool.Name)
	if runtime.GOOS == "windows" {
		candidatePath = candidatePath + ".exe"
	}
	if _, err := os.Stat(candidatePath); err == nil {
		return candidatePath, nil
	}

	// check binary is only file in src directory
	entries, err := os.ReadDir(srcDir)
	if err == nil && len(entries) == 1 && !entries[0].IsDir() {
		candidatePath = filepath.Join(srcDir, entries[0].Name())
		return candidatePath, nil
	}

	// check binary is named similar to the downloaded file name
	downloadedUrl, _ := tool.GetUrl() // ignore errs
	downloadedFileName := filepath.Base(downloadedUrl)
	if extIndex := strings.LastIndex(downloadedFileName, "."); extIndex > 0 {
		downloadedFileName = downloadedFileName[:strings.LastIndex(downloadedFileName, ".")]
	}
	candidatePath = filepath.Join(srcDir, downloadedFileName)
	if runtime.GOOS == "windows" {
		candidatePath = candidatePath + ".exe"
	}
	if _, err := os.Stat(candidatePath); err == nil {
		return candidatePath, nil
	}

	return "", ErrBinaryNotFound
}
