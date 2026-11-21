// Package files 设备侧文件系统(纯 Go 标准库)。路径为设备绝对路径;项目 resources 下用 Path() 取。
//
// 能力总览:
//   - 判断:Exists / IsFile / IsDir。
//   - 读写:Read / ReadBytes、Write / WriteBytes、Append 追加(均自动建父目录)。
//   - 目录:EnsureDir、ListDir 列一级、ClearDir 清空。
//   - 增删移:Copy / Move / Rename / Remove。
//   - 路径与摘要:Path 取 resources 相对路径、GetName / GetNameWithoutExtension / GetExtension、GetMd5。
package files

// Exists 路径是否存在(文件或目录)。
func Exists(path string) bool { var r0 bool; return r0 }

// IsFile 路径存在且是普通文件。
func IsFile(path string) bool { var r0 bool; return r0 }

// IsDir 路径存在且是目录。
func IsDir(path string) bool { var r0 bool; return r0 }

// EnsureDir 确保目录存在(递归)。
func EnsureDir(path string) error { var r0 error; return r0 }

// Read 读取整个文件为字符串。
func Read(path string) (string, error) { var r0 string; var r1 error; return r0, r1 }

// ReadBytes 读取整个文件为字节。
func ReadBytes(path string) ([]byte, error) { var r0 []byte; var r1 error; return r0, r1 }

// Write 写入文本(覆盖;父目录不存在则自动创建)。
func Write(path, text string) error { var r0 error; return r0 }

// WriteBytes 写入字节(覆盖;父目录不存在则自动创建)。
func WriteBytes(path string, data []byte) error { var r0 error; return r0 }

// Append 追加文本(不存在则创建)。
func Append(path, text string) error { var r0 error; return r0 }

// Copy 复制文件内容 from → to(覆盖;to 的父目录自动创建)。
func Copy(from, to string) error { var r0 error; return r0 }

// Move 移动 from → to(to 的父目录自动创建)。
func Move(from, to string) error { var r0 error; return r0 }

// Rename 在原目录内把 path 改名为 newName(仅文件名,不含目录)。
func Rename(path, newName string) error { var r0 error; return r0 }

// Remove 删除文件或目录(递归)。
func Remove(path string) error { var r0 error; return r0 }

// ClearDir 清空目录内容,返回删掉的条目数。
func ClearDir(path string) (int, error) { var r0 int; var r1 error; return r0, r1 }

// GetName 取路径的文件名部分(含扩展名)。
func GetName(path string) string { var r0 string; return r0 }

// GetNameWithoutExtension 取文件名并去掉扩展名。
func GetNameWithoutExtension(path string) string { var r0 string; return r0 }

// GetExtension 取扩展名(不含点,如 "png");无扩展名返回空串。
func GetExtension(path string) string { var r0 string; return r0 }

// GetMd5 文件 MD5(hex)。
func GetMd5(path string) (string, error) { var r0 string; var r1 error; return r0, r1 }

// ListDir 列目录(仅一级)。
func ListDir(path string) ([]Entry, error) { var r0 []Entry; var r1 error; return r0, r1 }
