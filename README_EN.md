# go-kit

A collection of common Go utility libraries providing file operations, string processing, data structures, time handling, encryption, and other common functionalities.

## Modules

### File Operations (file)
- File read/write, copy, move, delete
- Directory operations (create, delete, list, copy)
- File compression/decompression (tar, tar.gz)
- MD5 hash calculation

### String Processing (str)
- Basic string operations (substring, replace, reverse, etc.)
- Naming format conversion (camelCase, snake_case, kebab-case, etc.)
- String validation (email, URL, IP, phone number, etc.)

### Slice Operations (slice)
- Basic slice operations (unique, reverse, chunk, etc.)
- Filter and find operations
- Map and reduce transformations

### Data Structures (structs)
- Stack
- Queue
- Set
- LRU Cache

### Time Handling (time)
- Time formatting
- Time parsing
- Time calculations (add, subtract, diff, etc.)

### Crypto Tools (crypto)
- SHA series hashing (SHA1, SHA256, SHA512)
- AES encryption/decryption
- Base64 encoding/decoding

### Type Conversion (convert)
- String to number conversion
- String to boolean conversion
- Byte array to string conversion

### JSON Operations (json)
- JSON serialization and deserialization
- JSON pretty printing
- JSON field operations (get, set, merge)

### Random Number Generation (rand)
- Random integers and floats
- Random string generation
- Random element selection

### Database Operations (db)
- MySQL connection (GORM, XORM)
- Redis connection

### Network Operations (net)
- HTTP requests (GET, POST)
- IP address handling

## Installation

```bash
go get github.com/zzhuang94/go-kit
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
	"github.com/zzhuang94/go-kit/str"
	"github.com/zzhuang94/go-kit/slice"
)

func main() {
	// File operations
	file.WriteString("test.txt", "Hello, World!")
	content, _ := file.ReadString("test.txt")
	fmt.Println(content)
	
	// String processing
	fmt.Println(str.CamelCase("hello_world")) // "helloWorld"
	
	// Slice operations
	numbers := []int{1, 2, 3, 2, 4}
	unique := slice.Unique(numbers)
	fmt.Println(unique) // [1, 2, 3, 4]
}
```

## Documentation

Each module has detailed documentation. Please refer to the README.md files in each subdirectory.

## License

MIT License
