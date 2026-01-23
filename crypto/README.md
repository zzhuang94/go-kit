# Crypto 工具包

提供常用的加密和哈希工具函数。

## SHA 哈希

- `SHA1(data []byte) []byte` - 计算字节数组的 SHA1 值
- `SHA256(data []byte) []byte` - 计算字节数组的 SHA256 值
- `SHA512(data []byte) []byte` - 计算字节数组的 SHA512 值
- `SHA1String(s string) string` - 计算字符串的 SHA1 值
- `SHA256String(s string) string` - 计算字符串的 SHA256 值
- `SHA512String(s string) string` - 计算字符串的 SHA512 值
- `SHA1File(path string) (string, error)` - 计算文件的 SHA1 值
- `SHA256File(path string) (string, error)` - 计算文件的 SHA256 值
- `SHA512File(path string) (string, error)` - 计算文件的 SHA512 值

## AES 加密

- `AESEncrypt(key, plaintext []byte) ([]byte, error)` - AES加密
- `AESDecrypt(key, ciphertext []byte) ([]byte, error)` - AES解密
- `AESEncryptString(key, plaintext string) (string, error)` - AES加密（字符串）
- `AESDecryptString(key, ciphertext string) (string, error)` - AES解密（字符串）

**注意**：AES密钥长度必须是 16、24 或 32 字节。

## Base64 编码

- `Base64Encode(data []byte) []byte` - Base64编码
- `Base64Decode(data []byte) ([]byte, error)` - Base64解码
- `Base64EncodeString(s string) string` - Base64编码（字符串）
- `Base64DecodeString(s string) (string, error)` - Base64解码（字符串）

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/crypto"
)

func main() {
	// SHA 哈希
	hash := crypto.SHA256String("hello world")
	fmt.Println(hash)
	
	// AES 加密
	key := "1234567890123456"
	encrypted, _ := crypto.AESEncryptString(key, "secret message")
	fmt.Println(encrypted)
	decrypted, _ := crypto.AESDecryptString(key, encrypted)
	fmt.Println(decrypted)
	
	// Base64 编码
	encoded := crypto.Base64EncodeString("hello world")
	fmt.Println(encoded)
	decoded, _ := crypto.Base64DecodeString(encoded)
	fmt.Println(decoded)
}
```

## 运行测试

```bash
go test ./crypto -v
```
