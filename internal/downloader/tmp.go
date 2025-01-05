package downloader

import (
	"os"
)

func CreateTmpDir() (string, error) {
	return os.MkdirTemp("./tmp", "glug-*") // todo: get tmp parent dir from config
}

func CleanUpTmpDir(dir string) {
	os.RemoveAll(dir)
}
