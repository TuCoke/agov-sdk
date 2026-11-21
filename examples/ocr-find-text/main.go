package main

import (
	"fmt"

	"github.com/TuCoke/agov-sdk/motion"
	"github.com/TuCoke/agov-sdk/ocr"
	"github.com/TuCoke/agov-sdk/utils"
)

// OCR 找字 -> 点击文字中心。
func main() {
	if err := ocr.Load("resources/ocr"); err != nil {
		utils.LogE("ocr", err.Error())
		return
	}
	defer ocr.Close()

	r, ok := ocr.FindText("确定")
	if !ok {
		utils.Log("没找到文字")
		return
	}
	motion.Tap(r.X+r.W/2, r.Y+r.H/2)
	fmt.Println("点击:", r.Text, r.Score)
}
