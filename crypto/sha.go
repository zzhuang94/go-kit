package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

// SHA1 计算字节数组的 SHA1 值 / Calculate SHA1 hash of byte array
func SHA1(data []byte) []byte {
	hash := sha1.Sum(data)
	return hash[:]
}

// SHA256 计算字节数组的 SHA256 值 / Calculate SHA256 hash of byte array
func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// SHA512 计算字节数组的 SHA512 值 / Calculate SHA512 hash of byte array
func SHA512(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}

// SHA1String 计算字符串的 SHA1 值 / Calculate SHA1 hash of string
func SHA1String(s string) string {
	hash := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// SHA256String 计算字符串的 SHA256 值 / Calculate SHA256 hash of string
func SHA256String(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// SHA512String 计算字符串的 SHA512 值 / Calculate SHA512 hash of string
func SHA512String(s string) string {
	hash := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// SHA1File 计算文件的 SHA1 值 / Calculate SHA1 hash of file
func SHA1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := sha1.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// SHA256File 计算文件的 SHA256 值 / Calculate SHA256 hash of file
func SHA256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// SHA512File 计算文件的 SHA512 值 / Calculate SHA512 hash of file
func SHA512File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := sha512.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
