//agov:shared
package ocr

// Result 一条识别结果。CX/CY 为归一化中心(0..1),X/Y/W/H 为像素包围盒。
type Result struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
	X     int     `json:"x"`
	Y     int     `json:"y"`
	W     int     `json:"w"`
	H     int     `json:"h"`
	CX    float64 `json:"cx"`
	CY    float64 `json:"cy"`
}
