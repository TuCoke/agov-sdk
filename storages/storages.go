// Package storages 本地 KV 存储(按表名分 JSON 文件,落在设备上,脚本重启后仍在)。纯 Go。
//
// 能力总览:Get / Put / Remove / Contains 按表读写单键、GetAll 取整表、Clear 清整表。
// 按 table 名分文件,适合存进度、账号态等需要跨脚本重启保留的小数据。并发安全。
package storages

import (
	"regexp"
	"sync"
)

var (
	baseDir = "/data/local/tmp/autogov-storage"
	mu      sync.Mutex
	safe    = regexp.MustCompile(`[^A-Za-z0-9_.-]`)
)

// Get 读取 table 表里 key 的值;不存在返回空串。
func Get(table, key string) string { var r0 string; return r0 }

// Put 在 table 表里写入 key=value(表文件不存在则创建)。
func Put(table, key, value string) error { var r0 error; return r0 }

// Remove 从 table 表里删除 key。
func Remove(table, key string) error { var r0 error; return r0 }

// Contains table 表里是否存在 key。
func Contains(table, key string) bool { var r0 bool; return r0 }

// GetAll 返回 table 表的全部键值对副本。
func GetAll(table string) map[string]string { var r0 map[string]string; return r0 }

// Clear 清空整张 table 表(删除其表文件)。
func Clear(table string) error { var r0 error; return r0 }
