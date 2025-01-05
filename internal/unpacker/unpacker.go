package unpacker

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	H_GZIP_TYPE = []byte{0x1F, 0x8B}
	H_ZIP_TYPE  = []byte{0x50, 0x4B, 0x03, 0x04}
)

func Unpack(src, dst string) (string, error) {
	srcReader, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to read package: %w", err)
	}
	defer srcReader.Close()

	header := make([]byte, 4)
	if _, err := io.ReadFull(srcReader, header); err != nil {
		return "", fmt.Errorf("unable to determine package file type: %w", err)
	}

	switch {
	case bytes.HasPrefix(header, H_GZIP_TYPE):
		return unpackGz(src, dst)
	case bytes.HasPrefix(header, H_ZIP_TYPE):
		return unpackZip(src, dst)
	default:
		return filepath.Dir(src), nil
	}
}
