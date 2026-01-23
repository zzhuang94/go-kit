# File 工具包

提供常用的文件和目录操作工具函数。

## 文件操作

### 文件存在性检查
- `Exists(path string) (bool, error)` - 检查文件或目录是否存在

### 文件读取
- `ReadAll(path string) ([]byte, error)` - 读取整个文件内容，返回字节数组
- `ReadString(path string) (string, error)` - 读取整个文件内容，返回字符串
- `ReadLines(path string) ([]string, error)` - 读取文件的所有行，返回字符串切片

### 文件写入（覆盖模式）
- `Write(path string, data []byte) error` - 写入字节数组到文件
- `WriteString(path string, content string) error` - 写入字符串到文件
- `WriteLines(path string, lines []string) error` - 写入多行字符串到文件

### 文件追加
- `Append(path string, data []byte) error` - 追加字节数组到文件末尾
- `AppendString(path string, content string) error` - 追加字符串到文件末尾
- `AppendLines(path string, lines []string) error` - 追加多行字符串到文件末尾

### 文件操作
- `Copy(src, dst string) error` - 复制文件到目标位置
- `Move(src, dst string) error` - 移动/重命名文件
- `Remove(path string) error` - 删除文件
- `Create(path string) error` - 创建空文件
- `Touch(path string) error` - 更新文件访问和修改时间为当前时间

### 文件信息
- `GetSize(path string) (int64, error)` - 获取文件大小（字节）
- `GetModTime(path string) (time.Time, error)` - 获取文件修改时间

### 文件权限检查
- `IsReadable(path string) bool` - 检查文件是否可读
- `IsWritable(path string) bool` - 检查文件是否可写

## 目录操作

### 目录检查
- `IsDir(path string) bool` - 判断指定路径是否为目录

### 目录创建
- `EnsureDir(dir string) error` - 确保目录存在，如果不存在则创建（包括所有父目录）

### 目录删除
- `RemoveDir(dir string) error` - 删除目录及其所有内容（递归删除）
- `CleanDir(dir string) error` - 清空目录内容但保留目录本身

### 目录移动
- `MoveDir(src, dst string) error` - 移动/重命名目录

### 目录列表
- `ListDir(dir string) ([]string, error)` - 列出目录下的所有文件和子目录名称
- `ListFiles(dir string) ([]string, error)` - 列出目录下的所有文件名称（不包括子目录）
- `ListDirs(dir string) ([]string, error)` - 列出目录下的所有子目录名称（不包括文件）

### 目录复制
- `CopyDir(src, dst string) error` - 复制目录及其所有内容到目标位置

### 目录信息
- `GetDirSize(dir string) (int64, error)` - 获取目录的总大小（字节），递归计算所有文件
- `GetDirModTime(dir string) (time.Time, error)` - 获取目录修改时间
- `IsDirEmpty(dir string) bool` - 检查目录是否为空
- `GetDirCount(dir string) (fileCount, dirCount int, err error)` - 获取目录下的文件和子目录数量

### 目录遍历
- `WalkDir(root string, walkFn filepath.WalkFunc) error` - 遍历目录，对每个文件或目录执行指定的函数

## MD5 操作

- `Md5(path string) (string, error)` - 计算文件的 MD5 值

## Tar 操作

### 打包
- `TarDir(srcDir, dstFile string) error` - 将目录打包为 tar 文件
- `TarGzDir(srcDir, dstFile string) error` - 将目录打包为 tar.gz 压缩文件

### 解压
- `Untar(srcFile, dstDir string) error` - 解压 tar 文件到指定目录
- `UntarGz(srcFile, dstDir string) error` - 解压 tar.gz 压缩文件到指定目录

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
)

func main() {
	// 写入文件
	if err := file.WriteString("test.txt", "Hello, World!"); err != nil {
		panic(err)
	}

	// 读取文件
	content, err := file.ReadString("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// 检查文件是否存在
	exists, err := file.Exists("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File exists:", exists)

	// 创建目录
	if err := file.EnsureDir("mydir"); err != nil {
		panic(err)
	}

	// 列出目录内容
	files, err := file.ListFiles("mydir")
	if err != nil {
		panic(err)
	}
	fmt.Println("Files:", files)

	// 计算文件 MD5
	hash, err := file.Md5("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("MD5:", hash)

	// 打包目录
	if err := file.TarGzDir("mydir", "mydir.tar.gz"); err != nil {
		panic(err)
	}

	// 解压文件
	if err := file.UntarGz("mydir.tar.gz", "extracted"); err != nil {
		panic(err)
	}
}
```

## 运行测试

```bash
go test ./file -v
```

测试会自动创建临时文件和目录，测试结束后会自动清理。
