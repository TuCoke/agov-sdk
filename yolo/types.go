//agov:shared
package yolo

// Result 一个检测框(屏幕像素)。
type Result struct {
	X, Y, Width, Height int
	Label               string
	Score               float64
	CenterX, CenterY    int
}
