package file

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

// IsDir 判断指定路径是否为目录 / Check if path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// EnsureDir 确保目录存在，如果不存在则创建（包括所有父目录）/ Ensure directory exists, create if not (including all parent directories)
func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// RemoveDir 删除目录及其所有内容（递归删除）/ Remove directory and all its contents (recursive)
func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}

// CleanDir 清空目录内容但保留目录本身 / Clean directory contents but keep directory itself
func CleanDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		} else {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}
	return nil
}

// MoveDir 移动/重命名目录 / Move or rename directory
func MoveDir(src, dst string) error {
	return os.Rename(src, dst)
}

// ListDir 列出目录下的所有文件和子目录名称 / List all files and subdirectories in directory
func ListDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}

// ListFiles 列出目录下的所有文件名称（不包括子目录）/ List all file names in directory (excluding subdirectories)
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := make([]string, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

// ListDirs 列出目录下的所有子目录名称（不包括文件）/ List all subdirectory names in directory (excluding files)
func ListDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	dirs := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

// CopyDir 复制目录及其所有内容到目标位置 / Copy directory and all its contents to destination
func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return os.ErrInvalid
	}
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		} else {
			return copyFile(path, dstPath, info.Mode())
		}
	})
}

func copyFile(src, dst string, mode os.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

// GetDirSize 获取目录的总大小（字节），递归计算所有文件 / Get total directory size in bytes, recursively calculate all files
func GetDirSize(dir string) (int64, error) {
	var size int64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// GetDirModTime 获取目录修改时间 / Get directory modification time
func GetDirModTime(dir string) (time.Time, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// IsDirEmpty 检查目录是否为空 / Check if directory is empty
func IsDirEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(entries) == 0
}

// GetDirCount 获取目录下的文件和子目录数量 / Get count of files and subdirectories in directory
func GetDirCount(dir string) (fileCount, dirCount int, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, 0, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
	}
	return fileCount, dirCount, nil
}

// WalkDir 遍历目录，对每个文件或目录执行指定的函数 / Walk directory, execute function for each file or directory
func WalkDir(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}
