# Crypto Utilities Package

Provides common encryption and hashing utility functions.

## SHA Hashing

- `SHA1(data []byte) []byte` - Calculate SHA1 hash of byte array
- `SHA256(data []byte) []byte` - Calculate SHA256 hash of byte array
- `SHA512(data []byte) []byte` - Calculate SHA512 hash of byte array
- `SHA1String(s string) string` - Calculate SHA1 hash of string
- `SHA256String(s string) string` - Calculate SHA256 hash of string
- `SHA512String(s string) string` - Calculate SHA512 hash of string
- `SHA1File(path string) (string, error)` - Calculate SHA1 hash of file
- `SHA256File(path string) (string, error)` - Calculate SHA256 hash of file
- `SHA512File(path string) (string, error)` - Calculate SHA512 hash of file

## AES Encryption

- `AESEncrypt(key, plaintext []byte) ([]byte, error)` - AES encryption
- `AESDecrypt(key, ciphertext []byte) ([]byte, error)` - AES decryption
- `AESEncryptString(key, plaintext string) (string, error)` - AES encryption (string)
- `AESDecryptString(key, ciphertext string) (string, error)` - AES decryption (string)

**Note**: AES key length must be 16, 24, or 32 bytes.

## Base64 Encoding

- `Base64Encode(data []byte) []byte` - Base64 encoding
- `Base64Decode(data []byte) ([]byte, error)` - Base64 decoding
- `Base64EncodeString(s string) string` - Base64 encoding (string)
- `Base64DecodeString(s string) (string, error)` - Base64 decoding (string)

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/crypto"
)

func main() {
	// SHA hashing
	hash := crypto.SHA256String("hello world")
	fmt.Println(hash)
	
	// AES encryption
	key := "1234567890123456"
	encrypted, _ := crypto.AESEncryptString(key, "secret message")
	fmt.Println(encrypted)
	decrypted, _ := crypto.AESDecryptString(key, encrypted)
	fmt.Println(decrypted)
	
	// Base64 encoding
	encoded := crypto.Base64EncodeString("hello world")
	fmt.Println(encoded)
	decoded, _ := crypto.Base64DecodeString(encoded)
	fmt.Println(decoded)
}
```

## Running Tests

```bash
go test ./crypto -v
```
