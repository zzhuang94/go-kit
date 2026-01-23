package file

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Md5 计算文件的 MD5 值 / Calculate MD5 hash of file
func Md5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil)), nil
}
