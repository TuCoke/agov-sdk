// Package images 截图、找色、找图。坐标为屏幕像素;区域 (x1,y1,x2,y2) 中 x2/y2<=0 表示到屏幕边缘。
//
// 颜色:"RRGGBB";偏色 "RRGGBB-DRDGDB";多选 "A|B"。多点找色见各函数。
// sim/thresh 0..1,越大越严。找不到返回 (-1,-1) 或 Match.Found=false。
//
// 实现是【胶水】:Go 抓帧(优先桥的 in-process 截屏,回退 screencap)+ 解析颜色串,
// 像素循环与模板匹配在 libagovkit.so(闭源二进制)里做。
package images

import (
	"image"
)

// grabRGB 取一帧紧凑 RGB:优先桥 in-process(快),回退 screencap。

// Screencap 全屏截图 JPEG。targetW 目标宽(<=0 原始),quality 1..100。
func Screencap(targetW, quality int) []byte { var r0 []byte; return r0 }

// Capture 区域截图为 *image.NRGBA(x2/y2<=0 到边);可再用 images 的图像处理函数加工。
func Capture(x1, y1, x2, y2 int) *image.NRGBA { var r0 *image.NRGBA; return r0 }

// Pixel 取 (x,y) 像素颜色 "RRGGBB";越界返回 ""。
func Pixel(x, y int) string { var r0 string; return r0 }

// CmpColor 判断 (x,y) 是否与 color 相似。
func CmpColor(x, y int, color string, sim float64) bool { var r0 bool; return r0 }

// FindColor 区域内找第一个相似像素,返回像素坐标;找不到 (-1,-1)。
func FindColor(x1, y1, x2, y2 int, color string, sim float64) (int, int) {
	var r0 int
	var r1 int
	return r0, r1
}

// CountColor 统计区域内相似像素数。
func CountColor(x1, y1, x2, y2 int, color string, sim float64) int { var r0 int; return r0 }

// DetectsMultiColors 绝对多点比色:所有点都命中各自颜色才为 true。
func DetectsMultiColors(colors string, sim float64) bool { var r0 bool; return r0 }

// FindMultiColors 相对多点找色:找锚点且各偏移点命中 → 返回锚点像素坐标;找不到 (-1,-1)。
func FindMultiColors(x1, y1, x2, y2 int, colors string, sim float64) (int, int) {
	var r0 int
	var r1 int
	return r0, r1
}

// FindMultiColorsAll 相对多点找色找全部,返回所有命中锚点。
func FindMultiColorsAll(x1, y1, x2, y2 int, colors string, sim float64) [][2]int {
	var r0 [][2]int
	return r0
}

// decodeTpl 解码模板图(PNG/JPEG)为紧凑灰度缓冲 + 宽高。

// FindImage 全屏找模板图(PNG/JPEG 字节),thresh 常用 0.8。
func FindImage(tpl []byte, thresh float64) Match { var r0 Match; return r0 }

// FindImageScaled 多尺度找图(模板与屏上尺寸不一致时,慢一些)。
func FindImageScaled(tpl []byte, thresh float64) Match { var r0 Match; return r0 }

// WaitImage 反复找图直到命中或超时(timeoutMs)。
func WaitImage(tpl []byte, thresh float64, timeoutMs int) Match { var r0 Match; return r0 }
