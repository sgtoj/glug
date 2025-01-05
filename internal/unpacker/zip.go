package unpacker

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unpackZip(src, dst string) (string, error) {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return "", fmt.Errorf("unable to open zip: %w", err)
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(dst, f.Name)

		// protect against zipslip vulnerability
		if !isSafePath(fpath, dst) {
			return "", fmt.Errorf("illegal file path in package: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return "", fmt.Errorf("cannot create directory from package: %w", err)
		}

		dstFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", fmt.Errorf("cannot create file from package: %w", err)
		}

		pkgFile, err := f.Open()
		if err != nil {
			dstFile.Close()
			return "", fmt.Errorf("cannot read file from package: %w", err)
		}

		if _, err := io.Copy(dstFile, pkgFile); err != nil {
			dstFile.Close()
			pkgFile.Close()
			return "", fmt.Errorf("cannot write file from package: %w", err)
		}

		dstFile.Close()
		pkgFile.Close()
	}

	return dst, nil
}

// isSafePath ensures file path  in package does not break out of destination directory
func isSafePath(path, dst string) bool {
	return strings.HasPrefix(path, filepath.Clean(dst)+string(os.PathSeparator))
}
