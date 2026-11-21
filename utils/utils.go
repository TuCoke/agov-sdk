//agov:shared

// Package utils 纯 Go 小工具:等待、随机、日志、类型转换。桌面和真机行为一致。
package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Sleep 等 ms 毫秒。
func Sleep(ms int) { time.Sleep(time.Duration(ms) * time.Millisecond) }

// SleepRange 等 [minMs, maxMs] 之间的随机毫秒(拟人间隔)。
func SleepRange(minMs, maxMs int) { Sleep(Random(minMs, maxMs)) }

// Random 取 [min, max] 闭区间随机整数。
func Random(min, max int) int {
	if max < min {
		min, max = max, min
	}
	return min + rand.Intn(max-min+1)
}

// Log 打一行日志(agov run 会实时显示)。
func Log(format string, a ...any) { log.Printf(format, a...) }

// LogI 打一行带标签的信息日志:[label] message。
func LogI(label, message string) { log.Printf("[%s] %s", label, message) }

// LogE 打一行带标签的错误日志:[label] ERROR message。
func LogE(label, message string) { log.Printf("[%s] ERROR %s", label, message) }

// Exit 结束脚本。
func Exit(code int) { os.Exit(code) }

// I2s int 转字符串。
func I2s(i int) string { return strconv.Itoa(i) }

// F2s float64 转字符串(不带多余尾零)。
func F2s(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }

// B2s bool 转字符串("true"/"false")。
func B2s(b bool) string { return strconv.FormatBool(b) }

// S2i 字符串转 int;无法解析返回 0。
func S2i(s string) int { i, _ := strconv.Atoi(s); return i }

// S2f 字符串转 float64;无法解析返回 0。
func S2f(s string) float64 { f, _ := strconv.ParseFloat(s, 64); return f }

// S2b 字符串转 bool;无法解析返回 false。
func S2b(s string) bool { b, _ := strconv.ParseBool(s); return b }

// Sprintf 等价 fmt.Sprintf(省一个 import)。
func Sprintf(format string, a ...any) string { return fmt.Sprintf(format, a...) }
