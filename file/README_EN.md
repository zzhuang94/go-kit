# File Utilities Package

Provides common file and directory operation utility functions.

## File Operations

### File Existence Check
- `Exists(path string) (bool, error)` - Check if file or directory exists

### File Reading
- `ReadAll(path string) ([]byte, error)` - Read entire file content, return byte array
- `ReadString(path string) (string, error)` - Read entire file content, return string
- `ReadLines(path string) ([]string, error)` - Read all lines from file, return string slice

### File Writing (Overwrite Mode)
- `Write(path string, data []byte) error` - Write byte array to file
- `WriteString(path string, content string) error` - Write string to file
- `WriteLines(path string, lines []string) error` - Write multiple lines to file

### File Appending
- `Append(path string, data []byte) error` - Append byte array to end of file
- `AppendString(path string, content string) error` - Append string to end of file
- `AppendLines(path string, lines []string) error` - Append multiple lines to end of file

### File Operations
- `Copy(src, dst string) error` - Copy file to destination
- `Move(src, dst string) error` - Move or rename file
- `Remove(path string) error` - Remove file
- `Create(path string) error` - Create empty file
- `Touch(path string) error` - Update file access and modification time to current time

### File Information
- `GetSize(path string) (int64, error)` - Get file size in bytes
- `GetModTime(path string) (time.Time, error)` - Get file modification time

### File Permission Check
- `IsReadable(path string) bool` - Check if file is readable
- `IsWritable(path string) bool` - Check if file is writable

## Directory Operations

### Directory Check
- `IsDir(path string) bool` - Check if path is a directory

### Directory Creation
- `EnsureDir(dir string) error` - Ensure directory exists, create if not (including all parent directories)

### Directory Deletion
- `RemoveDir(dir string) error` - Remove directory and all its contents (recursive)
- `CleanDir(dir string) error` - Clean directory contents but keep directory itself

### Directory Move
- `MoveDir(src, dst string) error` - Move or rename directory

### Directory Listing
- `ListDir(dir string) ([]string, error)` - List all files and subdirectories in directory
- `ListFiles(dir string) ([]string, error)` - List all file names in directory (excluding subdirectories)
- `ListDirs(dir string) ([]string, error)` - List all subdirectory names in directory (excluding files)

### Directory Copy
- `CopyDir(src, dst string) error` - Copy directory and all its contents to destination

### Directory Information
- `GetDirSize(dir string) (int64, error)` - Get total directory size in bytes, recursively calculate all files
- `GetDirModTime(dir string) (time.Time, error)` - Get directory modification time
- `IsDirEmpty(dir string) bool` - Check if directory is empty
- `GetDirCount(dir string) (fileCount, dirCount int, err error)` - Get count of files and subdirectories in directory

### Directory Traversal
- `WalkDir(root string, walkFn filepath.WalkFunc) error` - Walk directory, execute function for each file or directory

## MD5 Operations

- `Md5(path string) (string, error)` - Calculate MD5 hash of file

## Tar Operations

### Packing
- `TarDir(srcDir, dstFile string) error` - Pack directory into tar file
- `TarGzDir(srcDir, dstFile string) error` - Pack directory into tar.gz compressed file

### Extraction
- `Untar(srcFile, dstDir string) error` - Extract tar file to specified directory
- `UntarGz(srcFile, dstDir string) error` - Extract tar.gz compressed file to specified directory

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
)

func main() {
	// Write file
	if err := file.WriteString("test.txt", "Hello, World!"); err != nil {
		panic(err)
	}

	// Read file
	content, err := file.ReadString("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// Check if file exists
	exists, err := file.Exists("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File exists:", exists)

	// Create directory
	if err := file.EnsureDir("mydir"); err != nil {
		panic(err)
	}

	// List directory contents
	files, err := file.ListFiles("mydir")
	if err != nil {
		panic(err)
	}
	fmt.Println("Files:", files)

	// Calculate file MD5
	hash, err := file.Md5("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("MD5:", hash)

	// Pack directory
	if err := file.TarGzDir("mydir", "mydir.tar.gz"); err != nil {
		panic(err)
	}

	// Extract file
	if err := file.UntarGz("mydir.tar.gz", "extracted"); err != nil {
		panic(err)
	}
}
```

## Running Tests

```bash
go test ./file -v
```

Tests automatically create temporary files and directories, and clean them up after testing.
