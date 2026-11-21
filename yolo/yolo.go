// Package yolo 目标检测(YOLO,ncnn 推理)。cgo 薄封装,真身在 libyolo.so(二进制)。
// 用法:New(param, bin, labelsPath, targetSize, conf, nms) → Detect(区域)/DetectFromImage/Close。
// 模型为 ncnn 的 .param/.bin;放项目 resources/models/yolo/,用 files.Path 取路径。
package yolo

import (

	/*
	   #include <stdlib.h>
	   typedef struct { int x, y, w, h, cls; float score; } YoloBox;
	   int  yolo_load(const char* paramPath, const char* binPath, int targetSize);
	   int  yolo_detect(const unsigned char* rgb, int width, int height, float prob_thr, float nms_thr, YoloBox* out, int maxOut);
	   void yolo_close();
	*/

	"image"
)

// Yolo 一个已加载的检测器。
type Yolo struct {
	labels    []string
	conf, nms float32
	loaded    bool
}

// New 加载一个 ncnn YOLO 模型(param + bin)——用自己的模型就把这两个路径指到自己的文件。
//
//   - labelsPath:类别文件,每行一个类名;传空用内置 COCO80(yolov8n 预训练顺序)。
//   - targetSize:模型输入边长,yolov8 默认 640;按你的模型给。
//   - conf / nms:置信度阈值与 NMS 阈值,默认 0.25 / 0.45。
//
// 加载失败返回 error(带底层 rc 与 param 路径)。
//
//	y, err := yolo.New(
//		files.Path("models/yolo/best.param"),
//		files.Path("models/yolo/best.bin"),
//		files.Path("models/yolo/labels.txt"), // 空则用 COCO80
//		640, 0.25, 0.45)
func New(paramPath, binPath, labelsPath string, targetSize int, conf, nms float32) (*Yolo, error) {
	var r0 *Yolo
	var r1 error
	return r0, r1
}

// loadLabels 从文件按行读类别名(每行一个,忽略空行);读不到返回 nil。

// Detect 在屏幕像素区域内检测(x2/y2<=0 到边;坐标换算回全屏)。
func (y *Yolo) Detect(x1, y1, x2, y2 int) []Result { var r0 []Result; return r0 }

// DetectFromImage 从图像对象检测(不截屏)。
func (y *Yolo) DetectFromImage(img *image.NRGBA) []Result { var r0 []Result; return r0 }

// Close 释放模型。
func (y *Yolo) Close() {}
