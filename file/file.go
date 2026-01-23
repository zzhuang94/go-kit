package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Exists 检查文件或目录是否存在 / Check if file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadAll 读取整个文件内容，返回字节数组 / Read entire file content, return byte array
func ReadAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadString 读取整个文件内容，返回字符串 / Read entire file content, return string
func ReadString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines 读取文件的所有行，返回字符串切片 / Read all lines from file, return string slice
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ret = make([]string, 0, 1<<14)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ret = append(ret, line)
	}
	return ret, nil
}

// Write 写入字节数组到文件（覆盖模式）/ Write byte array to file (overwrite mode)
func Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// WriteString 写入字符串到文件（覆盖模式）/ Write string to file (overwrite mode)
func WriteString(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// WriteLines 写入多行字符串到文件（覆盖模式）/ Write multiple lines to file (overwrite mode)
func WriteLines(path string, lines []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	return writer.Flush()
}

// Append 追加字节数组到文件末尾 / Append byte array to end of file
func Append(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

// AppendString 追加字符串到文件末尾 / Append string to end of file
func AppendString(path string, content string) error {
	return Append(path, []byte(content))
}

// AppendLines 追加多行字符串到文件末尾 / Append multiple lines to end of file
func AppendLines(path string, lines []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	return writer.Flush()
}

// Copy 复制文件到目标位置 / Copy file to destination
func Copy(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

// Move 移动/重命名文件 / Move or rename file
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Remove 删除文件 / Remove file
func Remove(path string) error {
	return os.Remove(path)
}

// Create 创建空文件 / Create empty file
func Create(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}

// Touch 更新文件访问和修改时间为当前时间 / Update file access and modification time to current time
func Touch(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	now := time.Now()
	return os.Chtimes(path, now, now)
}

// GetSize 获取文件大小（字节）/ Get file size in bytes
func GetSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetModTime 获取文件修改时间 / Get file modification time
func GetModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// IsReadable 检查文件是否可读 / Check if file is readable
func IsReadable(path string) bool {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

// IsWritable 检查文件是否可写 / Check if file is writable
func IsWritable(path string) bool {
	if exists, _ := Exists(path); exists {
		file, err := os.OpenFile(path, os.O_WRONLY, 0)
		if err != nil {
			return false
		}
		file.Close()
		return true
	}
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		dir = "."
	}
	file, err := os.OpenFile(filepath.Join(dir, ".write_test"), os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		return false
	}
	file.Close()
	os.Remove(filepath.Join(dir, ".write_test"))
	return true
}
