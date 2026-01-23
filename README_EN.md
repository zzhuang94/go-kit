# 🚀 go-kit

> A feature-rich, production-ready Go utility library collection that makes development more efficient and code more elegant

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zzhuang94/go-kit?style=flat-square)](https://goreportcard.com/report/github.com/zzhuang94/go-kit)

**go-kit** is a carefully designed collection of Go utility libraries, providing developers with clean implementations of common functionalities. Whether you need file operations, string processing, data structures, time handling, or encryption/decryption, we have the utility functions you need. All functions are thoughtfully designed, support generics, provide comprehensive error handling, and come with detailed bilingual documentation.

## ✨ Features

- 🎯 **Ready to Use** - No complex configuration, import and use immediately
- 🔧 **Feature Rich** - Covers 11+ modules including file, string, data structures, time, encryption, and more
- 🚀 **High Performance** - Built on Go standard library with excellent performance
- 📚 **Well Documented** - Every function has bilingual comments and detailed documentation
- 🧪 **Well Tested** - Complete unit test coverage
- 🔒 **Type Safe** - Full utilization of Go generics for type safety
- 🌍 **Internationalized** - Bilingual documentation in Chinese and English

## 📦 Modules

### 📁 File Operations (file)
Powerful file and directory operation utilities supporting file read/write, copy, move, delete, directory creation, traversal, copying, and tar/tar.gz compression/decompression with MD5 hash calculation.

### 🔤 String Processing (str)
Rich string manipulation functions including basic operations (substring, replace, reverse, etc.), naming format conversion (camelCase, snake_case, kebab-case, etc.), and common validation (email, URL, IP, phone number, etc.).

### 📊 Slice Operations (slice)
Powerful slice operation utilities supporting deduplication, reversal, chunking, filtering, finding, mapping, reduction, and other functional programming operations, fully leveraging Go generics.

### 🏗️ Data Structures (structs)
Generic implementations of common data structures including Stack, Queue, Set, and LRU Cache with excellent performance and clean APIs.

### ⏰ Time Handling (time)
Convenient time operation utilities supporting formatting, parsing, calculations (add, subtract, diff, etc.), making time handling simple and intuitive.

### 🔐 Crypto Tools (crypto)
Complete encryption toolkit including SHA series hashing (SHA1, SHA256, SHA512), AES encryption/decryption, and Base64 encoding/decoding.

### 🔄 Type Conversion (convert)
Simple type conversion utilities supporting conversions between strings and numbers, booleans, and byte arrays.

### 📄 JSON Operations (json)
Convenient JSON processing utilities supporting serialization, deserialization, pretty printing, and field operations (get, set, merge).

### 🎲 Random Number Generation (rand)
Flexible random number generation utilities supporting random integers, floats, string generation, and random element selection.

### 💾 Database Operations (db)
Database connection utilities supporting MySQL (GORM, XORM) and Redis (standalone, cluster) connections.

### 🌐 Network Operations (net)
HTTP request utilities and IP address handling, simplifying network programming.

## 🚀 Quick Start

### Installation

```bash
go get github.com/zzhuang94/go-kit
```

### Usage Example

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
	"github.com/zzhuang94/go-kit/str"
	"github.com/zzhuang94/go-kit/slice"
	"github.com/zzhuang94/go-kit/time"
)

func main() {
	// File operations
	file.WriteString("test.txt", "Hello, World!")
	content, _ := file.ReadString("test.txt")
	fmt.Println(content)
	
	// String processing
	fmt.Println(str.CamelCase("hello_world"))  // "helloWorld"
	fmt.Println(str.IsEmail("test@example.com")) // true
	
	// Slice operations
	numbers := []int{1, 2, 3, 2, 4}
	unique := slice.Unique(numbers)
	fmt.Println(unique) // [1, 2, 3, 4]
	
	// Time handling
	now := time.Now()
	fmt.Println(time.FormatDateTime(now)) // "2024-01-15 10:30:45"
}
```

## 📚 Documentation

Each module has detailed documentation and usage examples. Please refer to the README.md files in each subdirectory:

- [File Operations (file)](./file/README_EN.md)
- [String Processing (str)](./str/README_EN.md)
- [Slice Operations (slice)](./slice/README_EN.md)
- [Data Structures (structs)](./structs/README_EN.md)
- [Time Handling (time)](./time/README_EN.md)
- [Crypto Tools (crypto)](./crypto/README_EN.md)
- [Type Conversion (convert)](./convert/README_EN.md)
- [JSON Operations (json)](./json/README_EN.md)
- [Random Number Generation (rand)](./rand/README_EN.md)
- [Database Operations (db)](./db/README_EN.md)
- [Network Operations (net)](./net/README_EN.md)

## 🧪 Running Tests

```bash
# Run all tests
make test

# Or use Go command
go test ./...

# Run tests with coverage
go test ./... -cover
```

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

⭐ If this project helps you, please give it a Star!
