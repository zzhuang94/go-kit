# 🚀 go-kit

> A feature-rich, production-ready Go utility library collection that makes development more efficient and code more elegant

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zzhuang94/go-kit?style=flat-square)](https://goreportcard.com/report/github.com/zzhuang94/go-kit)

**go-kit** is a carefully designed collection of Go utility libraries, providing developers with clean implementations of common functionalities. Whether you need file operations, string processing, data structures, time handling, or encryption/decryption, we have the utility functions you need. All functions are thoughtfully designed, support generics, provide comprehensive error handling, and come with detailed bilingual documentation.

## ✨ Features

- 🎯 **Ready to Use** - No complex configuration, import and use immediately
- 🔧 **Feature Rich** - Covers 9+ modules including file, string, data structures, time, encryption, and more
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

### 🔢 Algorithms & Data Structures (algo)
Generic implementations of common data structures and algorithms including Stack, Queue, Set, LRU Cache, sorting algorithms (quicksort), and search algorithms (binary search, linear search, etc.) with excellent performance and clean APIs.

### ⏰ Time Handling (time)
Convenient time operation utilities supporting formatting, parsing, calculations (add, subtract, diff, etc.), making time handling simple and intuitive.

### 🔐 Crypto Tools (crypto)
Complete encryption toolkit including SHA series hashing (SHA1, SHA256, SHA512), AES encryption/decryption, and Base64 encoding/decoding.

### 🛠️ General Utilities (lib)
Practical general utility collection including random operations (random selection, shuffle), Shell command execution, log management (based on logrus with log rotation), distributed locks (Redis-based), rate limiter (supports Redis and local mode), and distributed election (Redis-based).

### 💾 Database Operations (db)
Database connection utilities supporting MySQL (GORM, XORM) and Redis (standalone, cluster) connections.

### 🌐 Network Operations (net)
HTTP request utilities and IP address handling, simplifying network programming.

## 📋 Function List

### 📁 File Operations (file)
- `file.Exists(path)` - Check if file or directory exists
- `file.ReadAll(path)` - Read entire file content, return byte array
- `file.ReadString(path)` - Read entire file content, return string
- `file.ReadLines(path)` - Read all lines from file, return string slice
- `file.Write(path, data)` - Write byte array to file (overwrite mode)
- `file.WriteString(path, content)` - Write string to file (overwrite mode)
- `file.WriteLines(path, lines)` - Write multiple lines to file (overwrite mode)
- `file.Append(path, data)` - Append byte array to end of file
- `file.AppendString(path, content)` - Append string to end of file
- `file.AppendLines(path, lines)` - Append multiple lines to end of file
- `file.Copy(src, dst)` - Copy file to destination
- `file.Move(src, dst)` - Move or rename file
- `file.Remove(path)` - Remove file
- `file.Create(path)` - Create empty file
- `file.Touch(path)` - Update file access and modification time to current time
- `file.GetSize(path)` - Get file size in bytes
- `file.GetModTime(path)` - Get file modification time
- `file.IsReadable(path)` - Check if file is readable
- `file.IsWritable(path)` - Check if file is writable
- `file.IsDir(path)` - Check if path is a directory
- `file.EnsureDir(dir)` - Ensure directory exists, create if not (including all parent directories)
- `file.RemoveDir(dir)` - Remove directory and all its contents (recursive)
- `file.CleanDir(dir)` - Clean directory contents but keep directory itself
- `file.MoveDir(src, dst)` - Move or rename directory
- `file.ListDir(dir)` - List all files and subdirectories in directory
- `file.ListFiles(dir)` - List all file names in directory (excluding subdirectories)
- `file.ListDirs(dir)` - List all subdirectory names in directory (excluding files)
- `file.CopyDir(src, dst)` - Copy directory and all its contents to destination
- `file.GetDirSize(dir)` - Get total directory size in bytes, recursively calculate all files
- `file.GetDirModTime(dir)` - Get directory modification time
- `file.IsDirEmpty(dir)` - Check if directory is empty
- `file.GetDirCount(dir)` - Get count of files and subdirectories in directory
- `file.WalkDir(root, walkFn)` - Walk directory, execute function for each file or directory
- `file.Md5(path)` - Calculate MD5 hash of file
- `file.TarDir(srcDir, dstFile)` - Pack directory into tar file
- `file.TarGzDir(srcDir, dstFile)` - Pack directory into tar.gz compressed file
- `file.Untar(srcFile, dstDir)` - Extract tar file to specified directory
- `file.UntarGz(srcFile, dstDir)` - Extract tar.gz compressed file to specified directory

### 🔤 String Processing (str)
- `str.Trim(s)` - Trim whitespace from both ends of string
- `str.Md5(params)` - Calculate MD5 hash of params
- `str.Uuid()` - Generate UUID
- `str.Sub(s, start, end)` - Extract substring, supports negative indices
- `str.Reverse(s)` - Reverse string
- `str.PadLeft(s, length, pad)` - Pad string to specified length on the left
- `str.PadRight(s, length, pad)` - Pad string to specified length on the right
- `str.Truncate(s, length)` - Truncate string to specified length
- `str.SplitLines(str)` - Split string by lines, remove empty lines and duplicates
- `str.ParseAndFormatJson(str)` - Parse and format JSON string
- `str.CamelCase(s)` - Convert to camelCase
- `str.SnakeCase(s)` - Convert to snake_case
- `str.KebabCase(s)` - Convert to kebab-case
- `str.TitleCase(s)` - Convert to Title Case (first letter of each word capitalized)
- `str.Pluralize(s)` - Simple pluralization (basic rules)
- `str.IsEmail(s)` - Validate email format
- `str.IsURL(s)` - Validate URL format
- `str.IsIP(s)` - Validate IP address format (IPv4)
- `str.IsPhone(s)` - Validate phone number format (Mainland China)
- `str.IsNumeric(s)` - Validate if string is numeric
- `str.IsAlpha(s)` - Validate if string is alphabetic
- `str.IsAlphaNumeric(s)` - Validate if string is alphanumeric
- `str.IsEmpty(s)` - Check if string is empty
- `str.IsBlank(s)` - Check if string is blank (empty or only whitespace)

### 📊 Slice Operations (slice)
- `slice.Contains(slice, item)` - Check if slice contains element
- `slice.IndexOf(slice, item)` - Find element index in slice, return -1 if not found
- `slice.LastIndexOf(slice, item)` - Find element index from end, return -1 if not found
- `slice.Unique(slice)` - Remove duplicate elements
- `slice.Reverse(slice)` - Reverse slice
- `slice.Shuffle(slice)` - Randomly shuffle slice
- `slice.Chunk(slice, size)` - Split slice into chunks
- `slice.Flatten(slices)` - Flatten nested slices
- `slice.Intersect(slice1, slice2)` - Get intersection of two slices
- `slice.Union(slice1, slice2)` - Get union of two slices
- `slice.Diff(slice1, slice2)` - Get difference of two slices (elements in slice1 but not in slice2)
- `slice.Remove(slice, item)` - Remove all matching elements
- `slice.RemoveAt(slice, index)` - Remove element at index
- `slice.Insert(slice, index, item)` - Insert element at index
- `slice.First(slice)` - Get first element, return zero value if slice is empty
- `slice.Last(slice)` - Get last element, return zero value if slice is empty
- `slice.Take(slice, n)` - Take first N elements
- `slice.Drop(slice, n)` - Drop first N elements
- `slice.Filter(slice, fn)` - Filter slice, keep elements matching condition
- `slice.Find(slice, fn)` - Find first element matching condition, return zero value if not found
- `slice.FindIndex(slice, fn)` - Find first element index matching condition, return -1 if not found
- `slice.FindLast(slice, fn)` - Find last element matching condition, return zero value if not found
- `slice.FindLastIndex(slice, fn)` - Find last element index matching condition, return -1 if not found
- `slice.Every(slice, fn)` - Check if all elements match condition
- `slice.Some(slice, fn)` - Check if at least one element matches condition
- `slice.Count(slice, fn)` - Count elements matching condition
- `slice.GroupBy(slice, fn)` - Group by function
- `slice.Map(slice, fn)` - Map transform slice
- `slice.Reduce(slice, initial, fn)` - Reduce slice
- `slice.FlatMap(slice, fn)` - Flat map
- `slice.Partition(slice, fn)` - Partition slice into two parts: matching and non-matching

### 🔢 Algorithms & Data Structures (algo)
- `algo.NewStack[T]()` - Create new stack
- `stack.Push(item)` - Push item onto stack
- `stack.Pop()` - Pop item from stack
- `stack.Peek()` - Peek at top element
- `stack.Size()` - Return stack size
- `stack.IsEmpty()` - Check if stack is empty
- `stack.Clear()` - Clear stack
- `algo.NewQueue[T]()` - Create new queue
- `queue.Enqueue(item)` - Enqueue item
- `queue.Dequeue()` - Dequeue item
- `queue.Peek()` - Peek at front element
- `queue.Size()` - Return queue size
- `queue.IsEmpty()` - Check if queue is empty
- `queue.Clear()` - Clear queue
- `algo.NewSet[T]()` - Create new set
- `set.Add(item)` - Add element
- `set.Remove(item)` - Remove element
- `set.Contains(item)` - Check if contains element
- `set.Size()` - Return set size
- `set.IsEmpty()` - Check if set is empty
- `set.Clear()` - Clear set
- `set.ToSlice()` - Convert to slice
- `set.Union(other)` - Get union
- `set.Intersect(other)` - Get intersection
- `set.Diff(other)` - Get difference (elements in s but not in other)
- `algo.NewLRU[K, V](capacity)` - Create LRU cache
- `lru.Get(key)` - Get value
- `lru.Set(key, value)` - Set value
- `lru.Remove(key)` - Remove key-value pair
- `lru.Clear()` - Clear cache
- `lru.Size()` - Return current size
- `lru.Cap()` - Return capacity
- `algo.Sort(slice)` - Sort slice using quicksort
- `algo.SortFunc(slice, less)` - Sort slice using custom comparison function
- `algo.SortDesc(slice)` - Sort slice in descending order
- `algo.IsSorted(slice)` - Check if slice is sorted
- `algo.IsSortedFunc(slice, less)` - Check if slice is sorted using custom comparison function
- `algo.BinarySearch(slice, target)` - Binary search, return index of target, -1 if not found
- `algo.BinarySearchFunc(slice, target, cmp)` - Binary search with custom comparison function
- `algo.LinearSearch(slice, target)` - Linear search, return index of target, -1 if not found
- `algo.FindFirst(slice, fn)` - Find first element index matching condition, return -1 if not found
- `algo.FindLast(slice, fn)` - Find last element index matching condition, return -1 if not found
- `algo.LowerBound(slice, target)` - Find first element index >= target
- `algo.UpperBound(slice, target)` - Find first element index > target

### ⏰ Time Handling (time)
- `time.Format(t, layout)` - Format time to string
- `time.FormatDate(t)` - Format date (YYYY-MM-DD)
- `time.FormatTime(t)` - Format time (HH:mm:ss)
- `time.FormatDateTime(t)` - Format date and time (YYYY-MM-DD HH:mm:ss)
- `time.FormatTimestamp(t)` - Format timestamp
- `time.Parse(layout, value)` - Parse time string
- `time.ParseDate(value)` - Parse date string (YYYY-MM-DD)
- `time.ParseTime(value)` - Parse time string (HH:mm:ss)
- `time.ParseDateTime(value)` - Parse date and time string (YYYY-MM-DD HH:mm:ss)
- `time.AddDays(t, days)` - Add days
- `time.AddHours(t, hours)` - Add hours
- `time.AddMinutes(t, minutes)` - Add minutes
- `time.DiffDays(t1, t2)` - Calculate days difference
- `time.DiffHours(t1, t2)` - Calculate hours difference
- `time.DiffMinutes(t1, t2)` - Calculate minutes difference
- `time.IsToday(t)` - Check if time is today
- `time.IsYesterday(t)` - Check if time is yesterday
- `time.IsTomorrow(t)` - Check if time is tomorrow
- `time.StartOfDay(t)` - Get start of day
- `time.EndOfDay(t)` - Get end of day
- `time.StartOfWeek(t)` - Get start of week (Monday)
- `time.StartOfMonth(t)` - Get start of month
- `time.StartOfYear(t)` - Get start of year

### 🔐 Crypto Tools (crypto)
- `crypto.SHA1(data)` - Calculate SHA1 hash of byte array
- `crypto.SHA256(data)` - Calculate SHA256 hash of byte array
- `crypto.SHA512(data)` - Calculate SHA512 hash of byte array
- `crypto.SHA1String(s)` - Calculate SHA1 hash of string
- `crypto.SHA256String(s)` - Calculate SHA256 hash of string
- `crypto.SHA512String(s)` - Calculate SHA512 hash of string
- `crypto.SHA1File(path)` - Calculate SHA1 hash of file
- `crypto.SHA256File(path)` - Calculate SHA256 hash of file
- `crypto.SHA512File(path)` - Calculate SHA512 hash of file
- `crypto.AESEncrypt(key, plaintext)` - AES encryption
- `crypto.AESDecrypt(key, ciphertext)` - AES decryption
- `crypto.AESEncryptString(key, plaintext)` - AES encryption (string)
- `crypto.AESDecryptString(key, ciphertext)` - AES decryption (string)
- `crypto.Base64Encode(data)` - Base64 encoding
- `crypto.Base64Decode(data)` - Base64 decoding
- `crypto.Base64EncodeString(s)` - Base64 encoding (string)
- `crypto.Base64DecodeString(s)` - Base64 decoding (string)

### 🛠️ General Utilities (lib)
- `lib.Choice(slice)` - Randomly select an element from slice
- `lib.Shuffle(slice)` - Randomly shuffle slice
- `lib.RunCmd(cmd, timeout...)` - Execute Shell command with optional timeout
- `logCfg.InitLogrus()` - Initialize global logrus logger
- `logCfg.BuildLogger()` - Create new logrus Logger instance
- `lib.GetFormatter()` - Get custom log formatter
- `lib.TryLock(c, key, timeout)` - Try to acquire distributed lock
- `lock.Release()` - Release lock
- `lib.TryCheckIn(c, key, limit, timeout)` - Try to check in to rate limiter
- `limiter.Release()` - Release rate limiter resource
- `lib.NewElection(key, val, cmdable)` - Create election instance
- `election.IsMaster()` - Check if current instance is master
- `election.GetMaster()` - Get current master node value
- `election.Release()` - Release master status
- `election.Stop()` - Stop election and release resources

### 💾 Database Operations (db)
- `mysqlCfg.ConnGorm()` - Connect MySQL using GORM
- `mysqlCfg.ConnXorm()` - Connect MySQL using XORM
- `db.ConnRedis(cfg)` - Connect standalone Redis
- `db.ConnRedisCluster(cfg)` - Connect Redis cluster

### 🌐 Network Operations (net)
- `net.Post(url, data, headers, timeoutSecond)` - Send POST request
- `net.Get(url, timeoutSecond)` - Send GET request
- `net.GetWithHeaders(url, timeoutSecond, headers)` - Send GET request with headers
- `net.ClientIp(req)` - Get client IP address
- `net.LocalIp()` - Get local IP address

## 🚀 Quick Start

### Installation

Install the latest version:
```bash
go get github.com/zzhuang94/go-kit@latest
```

Install a specific version (recommended):
```bash
go get github.com/zzhuang94/go-kit@v1.0.0
```

If you encounter version issues, you can directly specify the version in your project's `go.mod` file:
```go
require (
    github.com/zzhuang94/go-kit v1.0.0
)
```
Then run `go mod tidy`.

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
- [Algorithms & Data Structures (algo)](./algo/README_EN.md)
- [Time Handling (time)](./time/README_EN.md)
- [Crypto Tools (crypto)](./crypto/README_EN.md)
- [General Utilities (lib)](./lib/README_EN.md)
- [Database Operations (db)](./db/README_EN.md)
- [Network Operations (net)](./net/README_EN.md)

## 🧪 Running Tests

```bash
# Run all tests (including unit tests and benchmarks)
make test

# Or use one-click test scripts
# Windows
test.bat

# Unix/Linux/Mac
./test.sh

# Or use Go command
go test ./...

# Run tests with coverage
go test ./... -cover

# Run benchmarks only
make test-bench
# Or
go test ./... -bench=. -benchmem
```

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

⭐ If this project helps you, please give it a Star!
