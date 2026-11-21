//agov:shared
package images

// Match 一次找图结果:X/Y 命中中心像素,Score 相似度。
type Match struct {
	Found bool    `json:"found"`
	X     int     `json:"x"`
	Y     int     `json:"y"`
	Score float64 `json:"score"`
}
