package unpacker

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unpackGz(src, dst string) (string, error) {
	srcRaw, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to open package: %w", err)
	}
	defer srcRaw.Close()

	gzReader, err := gzip.NewReader(srcRaw)
	if err != nil {
		return "", fmt.Errorf("failed to open gzip reader for package: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read metadata for file from package: %w", err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filepath.Join(dst, header.Name), os.FileMode(header.Mode)); err != nil {
				return "", fmt.Errorf("failed to create directory from package: %w", err)
			}
		case tar.TypeReg:
			fpath := filepath.Join(dst, header.Name)
			if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return "", fmt.Errorf("failed to create directory from package: %w", err)
			}
			dstFile, err := os.Create(fpath)
			if err != nil {
				return "", fmt.Errorf("failed to create file from package: %w", err)
			}
			if _, err := io.Copy(dstFile, tarReader); err != nil {
				dstFile.Close()
				return "", fmt.Errorf("failed to write file from package: %w", err)
			}
		default:
			log.Println("skipping object in tar due to unknown type", header.Typeflag) // todo: convert to warn msg from logger
		}
	}

	return dst, nil
}
