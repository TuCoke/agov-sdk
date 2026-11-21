package main

import (
	"fmt"

	"github.com/TuCoke/agov-sdk/motion"
	"github.com/TuCoke/agov-sdk/uiacc"
)

// 控件树里找可点击节点并点它的中心。
func main() {
	nodes := uiacc.Dump()
	fmt.Println("节点数:", len(nodes))
	for _, n := range nodes {
		if n.Text == "登录" && n.Clickable {
			motion.Tap(n.CenterX(), n.CenterY())
			return
		}
	}
	fmt.Println("没找到目标节点")
}
