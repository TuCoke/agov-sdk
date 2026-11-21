package main

import (
	"fmt"

	"github.com/TuCoke/agov-sdk/images"
	"github.com/TuCoke/agov-sdk/motion"
	"github.com/TuCoke/agov-sdk/utils"
)

// 截图 -> 找色 -> 点击。
func main() {
	img := images.Capture(0, 0, 0, 0)
	if img == nil {
		utils.Log("截图失败")
		return
	}
	fmt.Println("屏幕尺寸:", img.Bounds())

	x, y := images.FindColor(0, 0, 0, 0, "FF0000", 0.9)
	if x < 0 || y < 0 {
		utils.Log("没找到红色目标")
		return
	}
	motion.Tap(x, y)
	utils.Sleep(500)
}
