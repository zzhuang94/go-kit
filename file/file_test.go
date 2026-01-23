package file

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var testDir string

func TestMain(m *testing.M) {
	var err error
	testDir, err = os.MkdirTemp("", "file_test_*")
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.RemoveAll(testDir)
	os.Exit(code)
}

func TestExists(t *testing.T) {
	testFile := filepath.Join(testDir, "test_exists.txt")
	exists, err := Exists(testFile)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("File should not exist")
	}
	WriteString(testFile, "test")
	exists, err = Exists(testFile)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Error("File should exist")
	}
}

func TestReadWrite(t *testing.T) {
	testFile := filepath.Join(testDir, "test_readwrite.txt")
	content := "Hello, World!"
	if err := WriteString(testFile, content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	read, err := ReadString(testFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if read != content {
		t.Errorf("Expected %q, got %q", content, read)
	}
	data := []byte("Binary data")
	if err := Write(testFile, data); err != nil {
		t.Fatalf("Write failed: %v", err)
	}
	readData, err := ReadAll(testFile)
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}
	if string(readData) != string(data) {
		t.Errorf("Expected %q, got %q", string(data), string(readData))
	}
}

func TestReadWriteLines(t *testing.T) {
	testFile := filepath.Join(testDir, "test_lines.txt")
	lines := []string{"line1", "line2", "line3"}
	if err := WriteLines(testFile, lines); err != nil {
		t.Fatalf("WriteLines failed: %v", err)
	}
	readLines, err := ReadLines(testFile)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}
	if len(readLines) != len(lines) {
		t.Errorf("Expected %d lines, got %d", len(lines), len(readLines))
	}
	for i, line := range lines {
		if readLines[i] != line {
			t.Errorf("Line %d: expected %q, got %q", i, line, readLines[i])
		}
	}
}

func TestAppend(t *testing.T) {
	testFile := filepath.Join(testDir, "test_append.txt")
	if err := WriteString(testFile, "Hello"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := AppendString(testFile, " World"); err != nil {
		t.Fatalf("AppendString failed: %v", err)
	}
	content, err := ReadString(testFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if content != "Hello World" {
		t.Errorf("Expected 'Hello World', got %q", content)
	}
	appendLines := []string{"line1", "line2"}
	if err := AppendLines(testFile, appendLines); err != nil {
		t.Fatalf("AppendLines failed: %v", err)
	}
	content2, err := ReadString(testFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if content2 != "Hello Worldline1\nline2\n" {
		t.Errorf("Expected 'Hello Worldline1\\nline2\\n', got %q", content2)
	}
	lines, err := ReadLines(testFile)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}
	if len(lines) < 2 {
		t.Errorf("Expected at least 2 lines, got %d", len(lines))
	}
}

func TestCopy(t *testing.T) {
	srcFile := filepath.Join(testDir, "src.txt")
	dstFile := filepath.Join(testDir, "dst.txt")
	content := "Copy test content"
	if err := WriteString(srcFile, content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := Copy(srcFile, dstFile); err != nil {
		t.Fatalf("Copy failed: %v", err)
	}
	copied, err := ReadString(dstFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if copied != content {
		t.Errorf("Expected %q, got %q", content, copied)
	}
}

func TestMove(t *testing.T) {
	srcFile := filepath.Join(testDir, "move_src.txt")
	dstFile := filepath.Join(testDir, "move_dst.txt")
	content := "Move test content"
	if err := WriteString(srcFile, content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := Move(srcFile, dstFile); err != nil {
		t.Fatalf("Move failed: %v", err)
	}
	if exists, _ := Exists(srcFile); exists {
		t.Error("Source file should not exist after move")
	}
	moved, err := ReadString(dstFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if moved != content {
		t.Errorf("Expected %q, got %q", content, moved)
	}
}

func TestCreate(t *testing.T) {
	testFile := filepath.Join(testDir, "test_create.txt")
	if err := Create(testFile); err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if exists, _ := Exists(testFile); !exists {
		t.Error("File should exist after Create")
	}
	size, err := GetSize(testFile)
	if err != nil {
		t.Fatalf("GetSize failed: %v", err)
	}
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}
}

func TestTouch(t *testing.T) {
	testFile := filepath.Join(testDir, "test_touch.txt")
	if err := Touch(testFile); err != nil {
		t.Fatalf("Touch failed: %v", err)
	}
	if exists, _ := Exists(testFile); !exists {
		t.Error("File should exist after Touch")
	}
	modTime, err := GetModTime(testFile)
	if err != nil {
		t.Fatalf("GetModTime failed: %v", err)
	}
	now := time.Now()
	diff := now.Sub(modTime)
	if diff < 0 {
		diff = -diff
	}
	if diff > 5*time.Second {
		t.Errorf("ModTime should be recent, got %v", modTime)
	}
}

func TestGetSize(t *testing.T) {
	testFile := filepath.Join(testDir, "test_size.txt")
	content := "12345"
	if err := WriteString(testFile, content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	size, err := GetSize(testFile)
	if err != nil {
		t.Fatalf("GetSize failed: %v", err)
	}
	if size != int64(len(content)) {
		t.Errorf("Expected size %d, got %d", len(content), size)
	}
}

func TestIsReadable(t *testing.T) {
	testFile := filepath.Join(testDir, "test_readable.txt")
	if err := WriteString(testFile, "test"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if !IsReadable(testFile) {
		t.Error("File should be readable")
	}
}

func TestIsWritable(t *testing.T) {
	testFile := filepath.Join(testDir, "test_writable.txt")
	if err := WriteString(testFile, "test"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if !IsWritable(testFile) {
		t.Error("File should be writable")
	}
	newFile := filepath.Join(testDir, "new_writable.txt")
	if !IsWritable(newFile) {
		t.Error("New file path should be writable")
	}
}

func TestIsDir(t *testing.T) {
	testDir := filepath.Join(testDir, "test_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if !IsDir(testDir) {
		t.Error("Should be a directory")
	}
	testFile := filepath.Join(testDir, "test.txt")
	if err := WriteString(testFile, "test"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if IsDir(testFile) {
		t.Error("Should not be a directory")
	}
}

func TestEnsureDir(t *testing.T) {
	testDir := filepath.Join(testDir, "ensure", "nested", "dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if !IsDir(testDir) {
		t.Error("Directory should exist")
	}
}

func TestRemoveDir(t *testing.T) {
	testDir := filepath.Join(testDir, "remove_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := RemoveDir(testDir); err != nil {
		t.Fatalf("RemoveDir failed: %v", err)
	}
	if exists, _ := Exists(testDir); exists {
		t.Error("Directory should not exist after RemoveDir")
	}
}

func TestCleanDir(t *testing.T) {
	testDir := filepath.Join(testDir, "clean_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file2.txt"), "test2"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := CleanDir(testDir); err != nil {
		t.Fatalf("CleanDir failed: %v", err)
	}
	if !IsDir(testDir) {
		t.Error("Directory should still exist")
	}
	if !IsDirEmpty(testDir) {
		t.Error("Directory should be empty")
	}
}

func TestListDir(t *testing.T) {
	testDir := filepath.Join(testDir, "list_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(testDir, "subdir")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	entries, err := ListDir(testDir)
	if err != nil {
		t.Fatalf("ListDir failed: %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}
}

func TestListFiles(t *testing.T) {
	testDir := filepath.Join(testDir, "list_files")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file2.txt"), "test2"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(testDir, "subdir")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	files, err := ListFiles(testDir)
	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestListDirs(t *testing.T) {
	testDir := filepath.Join(testDir, "list_dirs")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(testDir, "dir1")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(testDir, "dir2")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file.txt"), "test"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	dirs, err := ListDirs(testDir)
	if err != nil {
		t.Fatalf("ListDirs failed: %v", err)
	}
	if len(dirs) != 2 {
		t.Errorf("Expected 2 directories, got %d", len(dirs))
	}
}

func TestCopyDir(t *testing.T) {
	srcDir := filepath.Join(testDir, "copy_src")
	dstDir := filepath.Join(testDir, "copy_dst")
	if err := EnsureDir(srcDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(srcDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(srcDir, "subdir")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(srcDir, "subdir", "file2.txt"), "test2"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := CopyDir(srcDir, dstDir); err != nil {
		t.Fatalf("CopyDir failed: %v", err)
	}
	if !IsDir(dstDir) {
		t.Error("Destination directory should exist")
	}
	copied, err := ReadString(filepath.Join(dstDir, "file1.txt"))
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if copied != "test1" {
		t.Errorf("Expected 'test1', got %q", copied)
	}
}

func TestGetDirSize(t *testing.T) {
	testDir := filepath.Join(testDir, "size_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file1.txt"), "12345"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file2.txt"), "67890"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	size, err := GetDirSize(testDir)
	if err != nil {
		t.Fatalf("GetDirSize failed: %v", err)
	}
	if size != 10 {
		t.Errorf("Expected size 10, got %d", size)
	}
}

func TestIsDirEmpty(t *testing.T) {
	emptyDir := filepath.Join(testDir, "empty_dir")
	if err := EnsureDir(emptyDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if !IsDirEmpty(emptyDir) {
		t.Error("Directory should be empty")
	}
	if err := WriteString(filepath.Join(emptyDir, "file.txt"), "test"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if IsDirEmpty(emptyDir) {
		t.Error("Directory should not be empty")
	}
}

func TestGetDirCount(t *testing.T) {
	testDir := filepath.Join(testDir, "count_dir")
	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := WriteString(filepath.Join(testDir, "file2.txt"), "test2"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := EnsureDir(filepath.Join(testDir, "dir1")); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	fileCount, dirCount, err := GetDirCount(testDir)
	if err != nil {
		t.Fatalf("GetDirCount failed: %v", err)
	}
	if fileCount != 2 {
		t.Errorf("Expected 2 files, got %d", fileCount)
	}
	if dirCount != 1 {
		t.Errorf("Expected 1 directory, got %d", dirCount)
	}
}

func TestMd5(t *testing.T) {
	testFile := filepath.Join(testDir, "test_md5.txt")
	content := "Hello, MD5!"
	if err := WriteString(testFile, content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	hash, err := Md5(testFile)
	if err != nil {
		t.Fatalf("Md5 failed: %v", err)
	}
	if len(hash) != 32 {
		t.Errorf("Expected MD5 hash length 32, got %d", len(hash))
	}
	hash2, err := Md5(testFile)
	if err != nil {
		t.Fatalf("Md5 failed: %v", err)
	}
	if hash != hash2 {
		t.Error("MD5 hash should be consistent")
	}
}

func TestTarDir(t *testing.T) {
	srcDir := filepath.Join(testDir, "tar_src")
	dstFile := filepath.Join(testDir, "test.tar")
	if err := EnsureDir(srcDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(srcDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := WriteString(filepath.Join(srcDir, "file2.txt"), "test2"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := TarDir(srcDir, dstFile); err != nil {
		t.Fatalf("TarDir failed: %v", err)
	}
	if exists, _ := Exists(dstFile); !exists {
		t.Error("Tar file should exist")
	}
	extractDir := filepath.Join(testDir, "tar_extract")
	if err := Untar(dstFile, extractDir); err != nil {
		t.Fatalf("Untar failed: %v", err)
	}
	extracted, err := ReadString(filepath.Join(extractDir, "file1.txt"))
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if extracted != "test1" {
		t.Errorf("Expected 'test1', got %q", extracted)
	}
}

func TestTarGzDir(t *testing.T) {
	srcDir := filepath.Join(testDir, "targz_src")
	dstFile := filepath.Join(testDir, "test.tar.gz")
	if err := EnsureDir(srcDir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}
	if err := WriteString(filepath.Join(srcDir, "file1.txt"), "test1"); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if err := TarGzDir(srcDir, dstFile); err != nil {
		t.Fatalf("TarGzDir failed: %v", err)
	}
	if exists, _ := Exists(dstFile); !exists {
		t.Error("Tar.gz file should exist")
	}
	extractDir := filepath.Join(testDir, "targz_extract")
	if err := UntarGz(dstFile, extractDir); err != nil {
		t.Fatalf("UntarGz failed: %v", err)
	}
	extracted, err := ReadString(filepath.Join(extractDir, "file1.txt"))
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if extracted != "test1" {
		t.Errorf("Expected 'test1', got %q", extracted)
	}
}
