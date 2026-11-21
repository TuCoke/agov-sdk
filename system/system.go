// Package system 进程与 shell(胶水:shell + /proc)。
//
// 能力总览:
//   - Shell / ShellTimeout 执行任意 shell 命令(后者带毫秒超时)。
//   - 进程:Pid 按进程名/包名取 PID。
//   - 资源:CpuUsage 整机 CPU 占用、MemKB 某进程内存(KB)。
package system

// Shell 执行一条 shell 命令(sh -c),返回合并输出。
func Shell(cmd string) string { var r0 string; return r0 }

// ShellTimeout 带超时(毫秒)的 shell。
func ShellTimeout(cmd string, timeoutMs int) (string, error) {
	var r0 string
	var r1 error
	return r0, r1
}

// Pid 按进程名/包名取 PID(第一个);没有返回 0。
func Pid(name string) int { var r0 int; return r0 }

// MemKB 进程物理内存 KB(/proc/<pid>/status VmRSS)。
func MemKB(pid int) int { var r0 int; return r0 }

// CpuUsage 进程 CPU 占用百分比(top 单次采样,解析随 ROM 有差异)。
func CpuUsage(pid int) float64 { var r0 float64; return r0 }
