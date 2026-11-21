//agov:shared
package files

import (
	"os"
	"path/filepath"
)

// ResDir 项目 resources 目录在设备上的位置:取环境变量 AGOV_RES(agov run 会设),
// 没有则取可执行文件旁的 resources/。
func ResDir() string {
	if d := os.Getenv("AGOV_RES"); d != "" {
		return d
	}
	exe, err := os.Executable()
	if err != nil {
		return "resources"
	}
	return filepath.Join(filepath.Dir(exe), "resources")
}

// Path 把 resources 下的相对路径变成设备绝对路径,如 Path("tpl/ok.png")。
func Path(rel string) string { return filepath.Join(ResDir(), rel) }
