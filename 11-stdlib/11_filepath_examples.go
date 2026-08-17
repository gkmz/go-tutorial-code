package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// joinPath 拼接路径
func joinPath() {
	// filepath.Join 自动处理路径分隔符和斜杠
	// 自动处理不同操作系统的路径格式
	path := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Println("Joined path:", path)

	// 清理多余的斜杠和分隔符
	path2 := filepath.Join("home//user", "documents", "..", "file.txt")
	fmt.Println("With ..:", path2)
}

// getFileInfo 获取文件信息
func getFileInfo() {
	path := "/home/user/documents/report.pdf"

	// filepath.Abs 获取绝对路径
	absPath, _ := filepath.Abs(path)
	fmt.Println("Absolute path:", absPath)

	// filepath.Base 获取文件名
	base := filepath.Base(path)
	fmt.Println("Base name:", base)

	// filepath.Dir 获取目录
	dir := filepath.Dir(path)
	fmt.Println("Directory:", dir)

	// filepath.Ext 获取扩展名
	ext := filepath.Ext(path)
	fmt.Println("Extension:", ext)

	// filepath.Split 分割路径和文件名
	dir2, base2 := filepath.Split(path)
	fmt.Printf("Split - Dir: %q, Base: %q\n", dir2, base2)

	// filepath.VolumeName 获取卷名（Windows）
	// 在 Linux/Mac 上返回空字符串
	vol := filepath.VolumeName(path)
	fmt.Println("Volume name:", vol)
}

// cleanPath 清理路径
func cleanPath() {
	paths := []string{
		"/home/user/../user/./docs",
		"dir/subdir/../subdir/file.txt",
		"dir//subdir///file.txt",
	}

	for _, p := range paths {
		cleaned := filepath.Clean(p)
		fmt.Printf("Original: %s\n", p)
		fmt.Printf("Cleaned:  %s\n\n", cleaned)
	}
}

// checkPathType 判断路径类型
func checkPathType() {
	paths := []string{
		"/home/user/file.txt",
		"user/file.txt",
		"../user/file.txt",
		"file.txt",
	}

	for _, p := range paths {
		isAbs := filepath.IsAbs(p)
		fmt.Printf("IsAbs(%q) = %v\n", p, isAbs)
	}
}

// globMatch 路径匹配
func globMatch() {
	patterns := []string{"*.txt", "file-?.txt", "*.go", "test_*.go"}
	paths := []string{"report.txt", "file-1.txt", "main.go", "test_helper.go", "data.csv"}

	fmt.Println("Pattern matching:")
	for _, pattern := range patterns {
		fmt.Printf("\nPattern: %s\n", pattern)
		for _, path := range paths {
			match, err := filepath.Match(pattern, path)
			if err != nil {
				fmt.Printf("  Error matching %s: %v\n", path, err)
			} else if match {
				fmt.Printf("  %s matches\n", path)
			}
		}
	}
}

// relPath 计算相对路径
func relPath() {
	base := "/home/user/documents"
	target := "/home/user/documents/subdir/file.txt"

	// 计算从 base 到 target 的相对路径
	rel, err := filepath.Rel(base, target)
	if err != nil {
		fmt.Println("Rel error:", err)
	} else {
		fmt.Printf("Rel(%q, %q) = %q\n", base, target, rel)
	}

	// 另一个例子
	base2 := "/home/user"
	target2 := "/home/user/docs/file.txt"
	rel2, _ := filepath.Rel(base2, target2)
	fmt.Printf("Rel(%q, %q) = %q\n", base2, target2, rel2)
}

// walkDir 遍历目录
func walkDir() {
	// 模拟目录结构
	dirs := []string{
		"testdata",
		"testdata/subdir1",
		"testdata/subdir2",
		"testdata/subdir2/nested",
	}

	// 创建目录
	for _, dir := range dirs {
		fmt.Println("Would create:", dir)
	}

	// filepath.Walk 遍历目录
	// 注意：实际使用需要先创建目录
	fmt.Println("\nfilepath.Walk would iterate through:")
	fmt.Println("  testdata/")
	fmt.Println("  testdata/file1.txt")
	fmt.Println("  testdata/subdir1/")
	fmt.Println("  testdata/subdir1/file2.txt")
	fmt.Println("  testdata/subdir2/")
	fmt.Println("  testdata/subdir2/file3.txt")
	fmt.Println("  testdata/subdir2/nested/")
	fmt.Println("  testdata/subdir2/nested/file4.txt")
}

// fileExtensionExamples 文件扩展名处理
func fileExtensionExamples() {
	files := []string{
		"report.pdf",
		"image.jpg",
		"document",
		"archive.tar.gz",
		"file.txt.backup",
	}

	fmt.Println("File extension handling:")
	for _, f := range files {
		ext := filepath.Ext(f)
		base := strings.TrimSuffix(f, ext)
		fmt.Printf("  %-20s -> Base: %-15s Ext: %s\n", f, base, ext)
	}
}

func main() {
	fmt.Println("=== join path ===")
	joinPath()

	fmt.Println("\n=== get file info ===")
	getFileInfo()

	fmt.Println("\n=== clean path ===")
	cleanPath()

	fmt.Println("\n=== path type check ===")
	checkPathType()

	fmt.Println("\n=== glob matching ===")
	globMatch()

	fmt.Println("\n=== relative path ===")
	relPath()

	fmt.Println("\n=== walk directory ===")
	walkDir()

	fmt.Println("\n=== file extension ===")
	fileExtensionExamples()
}
