// Package ocr 端上文字识别(PP-OCRv5,ncnn 推理,不联网)。cgo 薄封装,真身在 libppocr.so(二进制)。
// 用前先 Load 模型目录(det.param/det.bin/rec.param/rec.bin);缺省用项目 resources/models/ocr。
package ocr

import (

	/*
	   #include <stdlib.h>
	   int   ppocr_load(const char* detParam, const char* detBin, const char* recParam, const char* recBin, int targetSize);
	   char* ppocr_detect(const unsigned char* rgb, int width, int height);
	   void  ppocr_free(char* p);
	   void  ppocr_close();
	*/

	"sync"
)

var (
	mu     sync.Mutex
	loaded bool
)

type raw struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
	X     int     `json:"x"`
	Y     int     `json:"y"`
	W     int     `json:"w"`
	H     int     `json:"h"`
}

// Load 加载模型目录;modelDir 为空用 resources/models/ocr;重复调用无副作用。
func Load(modelDir string) error { var r0 error; return r0 }

// Loaded 模型是否已加载。
func Loaded() bool { var r0 bool; return r0 }

// Region 识别像素区域内的文字(坐标换算回全屏);x2/y2<=0 到边。
func Region(x1, y1, x2, y2 int) []Result { var r0 []Result; return r0 }

// Screen 识别整屏文字。
func Screen() []Result { var r0 []Result; return r0 }

// FindText 整屏找包含 text 的第一条结果。
func FindText(text string) (Result, bool) { var r0 Result; var r1 bool; return r0, r1 }

// WaitText 反复识别直到出现包含 text 的文字或超时(timeoutMs)。
func WaitText(text string, timeoutMs int) (Result, bool) { var r0 Result; var r1 bool; return r0, r1 }

// Close 释放模型。
func Close() {}
